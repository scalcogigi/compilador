package parser

import (
	"compilador/lexer"
	"compilador/token"
)

type Parser struct {
	Lex *lexer.Lexer
}

// deve ser estático
func Run(code string) int {
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
func (p *Parser) parseExpression() int {

	result := p.parseTerm()

	for p.Lex.Next.Type == token.PLUS || p.Lex.Next.Type == token.MINUS {
		operator := p.Lex.Next.Type

		p.Lex.SelectNext()

		value := p.parseTerm()

		switch operator {
		case token.PLUS:
			result += value
		case token.MINUS:
			result -= value
		}

	}
	return result
}

// chama o número e depois chama o operador (* ou /) e o próximo número, até que não haja mais operadores
func (p *Parser) parseTerm() int {

	result := p.parseFactor()

	for p.Lex.Next.Type == token.MULT || p.Lex.Next.Type == token.DIV {
		operator := p.Lex.Next.Type

		p.Lex.SelectNext()

		value := p.parseFactor()

		switch operator {
		case token.MULT:
			result *= value
		case token.DIV:
			result /= value
		}
	}
	return result
}

// chama o número e depois chama o ( ou ) e o próximo número, até que não haja mais operadores
func (p *Parser) parseFactor() int {
	// CASO 1: número
	if p.Lex.Next.Type == token.INT {
		value := p.Lex.Next.Value.(int)
		
		p.Lex.SelectNext()

		return value
	} else if p.Lex.Next.Type == token.OPEN_PAREN {
		p.Lex.SelectNext() // consome '('

		result := p.parseExpression()

		if p.Lex.Next.Type != token.CLOSE_PAREN {
			panic("[Parser] Esperado ')', mas encontrado: " + p.Lex.Next.Type)
		}

		p.Lex.SelectNext() // consome ')'
		return result
	} else if p.Lex.Next.Type == token.PLUS || p.Lex.Next.Type == token.MINUS {
		operator := p.Lex.Next.Type
		p.Lex.SelectNext() // consome o operador

		value := p.parseFactor()

		if operator == token.MINUS {
			return -value
		}

		return value
	} else { panic("[Parser] Fator inválido: " + p.Lex.Next.Type)}
}