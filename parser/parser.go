package parser

import (
	"compilador/lexer"
	"compilador/node"
	"compilador/token"
)

type Parser struct {
	Lex *lexer.Lexer
}

func Run(code string) node.Node {
	parser := &Parser{
		Lex: lexer.New(code),
	}

	parser.Lex.SelectNext()
	result := parser.parseProgram()

	if parser.Lex.Next.Type != token.EOF {
		panic("[Parser] Esperado EOF, mas encontrado: " + parser.Lex.Next.Type)
	}

	return node.AssignIDs(result)
}

// parseProgram -> sequÊncia de instruções até EOF
func (p *Parser) parseProgram() node.Node {
	statements := []node.Node{}

	for p.Lex.Next.Type != token.EOF {
		statements = append(statements, p.parseStatement())
	}

	return node.Block{
		Children: statements,
	}
}

// parseBlock - sequência de instruções entre chaves
func (p *Parser) parseBlock() node.Node {
	if p.Lex.Next.Type != token.OPEN_BRA {
		panic("[Parser] Esperado '{', mas encontrado: " + p.Lex.Next.Type)
	}

	p.Lex.SelectNext() // consome token

	// p.parseEnd()

	statements := []node.Node{}

	for p.Lex.Next.Type != token.CLOSE_BRA && p.Lex.Next.Type != token.EOF {
		statements = append(statements, p.parseStatement())
	}

	if p.Lex.Next.Type != token.CLOSE_BRA {
		panic("[Parser] Esperado '}', mas encontrado: " + p.Lex.Next.Type)
	}
	p.Lex.SelectNext() // consome token

	return node.Block{
		Children: statements,
	}
}

// Valida e consome o fim de uma instrução.
func (p *Parser) parseEnd() {
	if p.Lex.Next.Type != token.END {
		panic("[Parser] Esperado fim de linha, mas encontrado: " + p.Lex.Next.Type)
	}
	p.Lex.SelectNext()
}

// parseStatement
func (p *Parser) parseStatement() node.Node {
	switch p.Lex.Next.Type {

	case token.END:
		p.Lex.SelectNext()
		return node.NoOp{}

	case token.OPEN_BRA:
		return p.parseBlock()

	// atribuição
	case token.IDEN:
		{
			identifier := p.Lex.Next.Value.(string)
			p.Lex.SelectNext()

			if p.Lex.Next.Type != token.ASSIGN {
				panic("[Parser] Esperado '=', mas encontrado: " + p.Lex.Next.Type)
			}
			p.Lex.SelectNext()

			expression := p.parseBoolExpression()

			p.parseEnd()

			return node.Assignment{
				Children: []node.Node{
					node.Identifier{Value: identifier},
					expression,
				},
			}
		}

	// condicional
	case token.IF:
		{
			p.Lex.SelectNext()

			value := p.parseBoolExpression()
			thenBlock := p.parseBlock()

			children := []node.Node{
				value,
				thenBlock,
			}

			if p.Lex.Next.Type == token.ELSE {
				p.Lex.SelectNext()

				elseBlock := p.parseBlock()
				children = append(children, elseBlock)
			}

			p.parseEnd()

			return node.IF{
				Children: children,
			}
		}

	// while
	case token.WHILE:
		{
			p.Lex.SelectNext()

			value := p.parseBoolExpression()
			body := p.parseBlock()

			p.parseEnd()

			return node.While{
				Children: []node.Node{
					value,
					body,
				},
			}
		}

	// print
	case token.PRINT:
		{
			p.Lex.SelectNext()

			if p.Lex.Next.Type != token.OPEN_PAREN {
				panic("[Parser] Esperado '(', mas encontrado: " + p.Lex.Next.Type)
			}
			p.Lex.SelectNext()

			expression := p.parseBoolExpression()

			if p.Lex.Next.Type != token.CLOSE_PAREN {
				panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
			}
			p.Lex.SelectNext()

			p.parseEnd()

			return node.Println{
				Children: []node.Node{
					expression,
				},
			}
		}

	// token VAR
	case token.VAR:
		p.Lex.SelectNext() // consome "var"

		if p.Lex.Next.Type != token.IDEN {
			panic("[Parser] Esperado identificador após 'var'")
		}
		name := p.Lex.Next.Value.(string)
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.TYPE {
			panic("[Parser] Esperado tipo na declaração")
		}
		declaredType := p.Lex.Next.Value.(string)
		p.Lex.SelectNext()

		children := []node.Node{
			node.Identifier{Value: name},
		}

		if p.Lex.Next.Type == token.ASSIGN {
			p.Lex.SelectNext()
			children = append(children, p.parseBoolExpression())
		}

		p.parseEnd()

		return node.VarDec{
			Value:    declaredType,
			Children: children,
		}
	}

	panic("[Parser] Instrução inválida: " + p.Lex.Next.Type)

}

// parseBoolExpression
func (p *Parser) parseBoolExpression() node.Node {
	result := p.parseBoolTerm()

	for p.Lex.Next.Type == token.OR {
		operator := p.Lex.Next.Type
		p.Lex.SelectNext()

		right := p.parseBoolTerm()

		result = node.BinOp{
			Value:    operator,
			Children: []node.Node{result, right},
		}
	}

	return result
}

// parseBoolTerm
func (p *Parser) parseBoolTerm() node.Node {
	result := p.parseRelExpression()

	for p.Lex.Next.Type == token.AND {
		operator := p.Lex.Next.Type
		p.Lex.SelectNext() // consome token

		right := p.parseRelExpression()

		result = node.BinOp{
			Value:    operator,
			Children: []node.Node{result, right},
		}
	}
	return result
}

// parseRelExpression
func (p *Parser) parseRelExpression() node.Node {
	result := p.parseExpression()

	if p.Lex.Next.Type == token.EQ ||
		p.Lex.Next.Type == token.LT ||
		p.Lex.Next.Type == token.GT {
		operator := p.Lex.Next.Type
		p.Lex.SelectNext()

		right := p.parseExpression()

		result = node.BinOp{
			Value:    operator,
			Children: []node.Node{result, right},
		}
	}
	return result
}

// chama o número e depois chama o operador(+ ou -) e o próximo número, até que não haja mais operadores
func (p *Parser) parseExpression() node.Node {

	result := p.parseTerm()

	for p.Lex.Next.Type == token.PLUS || p.Lex.Next.Type == token.MINUS {
		operator := p.Lex.Next.Type

		p.Lex.SelectNext()

		right := p.parseTerm()

		result = node.BinOp{
			Value:    operator,
			Children: []node.Node{result, right},
		}
	}
	return result
}

// chama o número e depois chama o operador (* ou /) e o próximo número, até que não haja mais operadores
func (p *Parser) parseTerm() node.Node {

	result := p.parseFactor()

	for p.Lex.Next.Type == token.MULT || p.Lex.Next.Type == token.DIV {
		operator := p.Lex.Next.Type

		p.Lex.SelectNext()

		value := p.parseFactor()

		result = node.BinOp{
			Value:    operator,
			Children: []node.Node{result, value},
		}
	}
	return result
}

// operadores unários (+ e -)
func (p *Parser) parseFactor() node.Node {

	if p.Lex.Next.Type == token.PLUS ||
		p.Lex.Next.Type == token.MINUS ||
		p.Lex.Next.Type == token.NOT {

		operator := p.Lex.Next.Type
		p.Lex.SelectNext() // consome token

		value := p.parseFactor()

		return node.UnOp{
			Value:    operator,
			Children: []node.Node{value},
		}
	}

	return p.parsePower()
}

func (p *Parser) parsePower() node.Node {
	base := p.parsePrimary()

	if p.Lex.Next.Type == token.POWER {
		p.Lex.SelectNext() // consome "**"

		exponent := p.parseFactor()

		return node.BinOp{
			Value:    token.POWER,
			Children: []node.Node{base, exponent},
		}
	}

	return base
}

func (p *Parser) parsePrimary() node.Node {

	if p.Lex.Next.Type == token.BOOL {
		value := p.Lex.Next.Value.(bool)
		p.Lex.SelectNext()
		return node.BoolVal{Value: value}
	}

	if p.Lex.Next.Type == token.STR {
		value := p.Lex.Next.Value.(string)
		p.Lex.SelectNext()
		return node.StringVal{Value: value}
	}

	if p.Lex.Next.Type == token.INT {
		value := p.Lex.Next.Value.(int)

		p.Lex.SelectNext()
		return node.IntVal{
			Value: value,
		}
	}

	if p.Lex.Next.Type == token.IDEN {
		identifier := p.Lex.Next.Value.(string)
		p.Lex.SelectNext()

		return node.Identifier{
			Value: identifier,
		}
	}

	if p.Lex.Next.Type == token.OPEN_PAREN {
		p.Lex.SelectNext() // consome "("

		result := p.parseBoolExpression()

		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}

		p.Lex.SelectNext() // consome ")"
		return result
	}

	// Scanln
	if p.Lex.Next.Type == token.READ {
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.OPEN_PAREN {
			panic("[Parser] Esperado '(', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		return node.ScanLn{}
	}

	panic("[Parser] Expressão primária inválida: " + p.Lex.Next.Type)
}
