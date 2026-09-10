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
	result := parser.parseExpression()

	if parser.Lex.Next.Type != token.EOF {
		panic("[Parser] Esperado EOF, mas encontrado: " + parser.Lex.Next.Type)
	}

	return result
}

// chama o número e depois chama o operador(+ ou -) e o próximo número, até que não haja mais operadores
func (p *Parser) parseExpression() node.Node {

	result := p.parseTerm()

	for p.Lex.Next.Type == token.PLUS || p.Lex.Next.Type == token.MINUS {
		operator := p.Lex.Next.Type

		p.Lex.SelectNext()

		value := p.parseTerm()

		result = node.BinOp{
			Value: operator,
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
			Value: operator,
			Children: []node.Node{result, value},
		}
	}
	return result
}

// operadores unários (+ e -)
func (p *Parser) parseFactor() node.Node {

	if p.Lex.Next.Type == token.PLUS ||
		p.Lex.Next.Type == token.MINUS {

		operator := p.Lex.Next.Type
		p.Lex.SelectNext() // consome "+" ou "-"

		value := p.parseFactor()

		return node.UnOp{
			Value: operator,
			Children: []node.Node{value},
		}
	}

	return p.parsePower()
}

func (p *Parser) parsePower() node.Node {
	base := p.parsePostfix()

	if p.Lex.Next.Type == token.POWER {
		p.Lex.SelectNext() // consome "**"

		exponent := p.parseFactor()

		return node.BinOp{
			Value: token.POWER,
			Children: []node.Node{base, exponent},
		}
	}

	return base
}

func (p *Parser) parsePostfix() node.Node {
	result := p.parsePrimary()

	for p.Lex.Next.Type == token.FACT {
		p.Lex.SelectNext() // consome "!"

		result = node.UnOp{
			Value:    token.FACT,
			Children: []node.Node{result},
		}
	}

	return result
}

func (p *Parser) parsePrimary() node.Node {

	if p.Lex.Next.Type == token.INT {
		value := p.Lex.Next.Value.(int)

		p.Lex.SelectNext()
		return node.IntVal{
			Value: value,
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