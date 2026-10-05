package parser

import (
	"compilador/lexer"
	"compilador/node"
	"compilador/token"
)

type Parser struct {
	Lex *lexer.Lexer
}

// deve ser estático
func Run(code string) node.Node {
	lex := lexer.New(code)

	parser := &Parser{
		Lex: lex,
	}

	parser.Lex.SelectNext()
	result := parser.parseProgram()

	if parser.Lex.Next.Type != token.EOF {
		panic("[Parser] Esperado EOF, mas encontrado: " + parser.Lex.Next.Type)
	}

	return result
}

// parseProgram
func (p *Parser) parseProgram() node.Node {
	statements := []node.Node{}

	for p.Lex.Next.Type != token.EOF {
		statement := p.parseStatement()
		statements = append(statements, statement)
	}

	return node.Block{
		Children: statements,
	}
}

// parseBlock
func (p *Parser) parseBlock() node.Node {
	if p.Lex.Next.Type != token.OPEN_BRA {
		panic("[Parser] Esperado '{', mas encontrado: " + p.Lex.Next.Type)
	}
	p.Lex.SelectNext() // consome token

	statements := []node.Node{}

	for p.Lex.Next.Type != token.CLOSE_BRA {
		statement := p.parseStatement()
		statements = append(statements, statement)
	}

	if p.Lex.Next.Type != token.CLOSE_BRA {
		panic("[Parser] Esperado '}', mas encontrado: " + p.Lex.Next.Type)
	}
	p.Lex.SelectNext() // consome token

	return node.Block{
		Children: statements,
	}
}

// parseStatement
func (p *Parser) parseStatement() node.Node {
	// vazio
	if p.Lex.Next.Type == token.END {
		p.Lex.SelectNext()
		return node.NoOp{}
	}

	// atribuição
	if p.Lex.Next.Type == token.IDEN {
		identifier := p.Lex.Next.Value.(string)
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.ASSIGN {
			panic("[Parser] Esperado '=', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		expression := p.parseExpression()

		if p.Lex.Next.Type != token.END {
			panic("[Parser] Esperado fim de linha, mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		return node.Assignment{
			Children: []node.Node{
				node.Identifier{Value: identifier},
				expression,
			},
		}
	}

	// condicional
	if p.Lex.Next.Type == token.IF {
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.OPEN_PAREN {
			panic("[Parser] Esperado '(', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		expression := p.parseExpression()

		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.END {
			panic("[Parser] Esperado fim de linha, mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		return node.IF{
			Children: []node.Node{
				expression,
			},
		}
	}

	// while
	if p.Lex.Next.Type == token.WHILE {
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.OPEN_PAREN {
			panic("[Parser] Esperado '(', mas encontrado: " + p.Lex.Next.Type)
		}	
		p.Lex.SelectNext()

		expression := p.parseExpression()
		
		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}

		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.END {
			panic("[Parser] Esperado fim de linha, mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		return node.While{
			Children: []node.Node{
				expression,
			},
		}
	}

	// Scanln
	if p.Lex.Next.Type == token.READ {
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.OPEN_PAREN {
			panic("[Parser] Esperado '(', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.IDEN {
			panic("[Parser] Esperado identificador, mas encontrado: " + p.Lex.Next.Type)
		}
		identifier := p.Lex.Next.Value.(string)
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.END {
			panic("[Parser] Esperado fim de linha, mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		return node.ScanLn{
			Children: []node.Node{
				node.Identifier{Value: identifier},
			},
		}
	}

	// print
	if p.Lex.Next.Type == token.PRINT {
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.OPEN_PAREN {
			panic("[Parser] Esperado '(', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		expression := p.parseExpression()

		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		if p.Lex.Next.Type != token.END {
			panic("[Parser] Esperado fim de linha, mas encontrado: " + p.Lex.Next.Type)
		}
		p.Lex.SelectNext()

		return node.Println{
			Children: []node.Node{
				expression,
			},
		}
	}

	panic("[Parser] Instrução inválida: " + p.Lex.Next.Type)
}

// chama o número e depois chama o operador(+ ou -) e o próximo número, até que não haja mais operadores
func (p *Parser) parseExpression() node.Node {

	result := p.parseTerm()

	for p.Lex.Next.Type == token.PLUS || p.Lex.Next.Type == token.MINUS {
		operator := p.Lex.Next.Type

		p.Lex.SelectNext()

		value := p.parseTerm()

		result = node.BinOp{
			Value:    operator,
			Children: []node.Node{result, value},
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

		result := p.parseExpression()

		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}

		p.Lex.SelectNext() // consome ")"
		return result
	}

	panic("[Parser] Expressão primária inválida: " + p.Lex.Next.Type)
}

// parseBoolExpression
func (p *Parser) parseBoolExpression() node.Node {
	left := p.parseExpression()

	if p.Lex.Next.Type == token.EQ ||
		p.Lex.Next.Type == token.LT ||
		p.Lex.Next.Type == token.GT {

		operator := p.Lex.Next.Type
		p.Lex.SelectNext() // consome token

		right := p.parseExpression()

		return node.BinOp{
			Value:    operator,
			Children: []node.Node{left, right},
		}
	}
	return left
}

// parseBoolTerm
func (p *Parser) parseBoolTerm() node.Node {
	left := p.parseBoolExpression()

	if p.Lex.Next.Type == token.AND {
		operator := p.Lex.Next.Type
		p.Lex.SelectNext() // consome token
		right := p.parseBoolExpression()

		return node.BinOp{
			Value:    operator,
			Children: []node.Node{left, right},
		}
	}	
	return left
}

// parseRelExpression
func (p *Parser) parseRelExpression() node.Node {
	left := p.parseBoolTerm()

	if p.Lex.Next.Type == token.OR {
		operator := p.Lex.Next.Type
		p.Lex.SelectNext() // consome token
		right := p.parseBoolTerm()

		return node.BinOp{
			Value:    operator,
			Children: []node.Node{left, right},
		}
	}
	return left
}