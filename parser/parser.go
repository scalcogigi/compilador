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

func (p *Parser) parseExpression() int {
	
	if p.Lex.Next.Type != token.INT {
		panic("[Parser] Esperado um número inteiro, mas encontrado: " + p.Lex.Next.Type)
	}

	result := p.Lex.Next.Value.(int)
	p.Lex.SelectNext()

	for p.Lex.Next.Type == token.PLUS || p.Lex.Next.Type == token.MINUS {
		operator := p.Lex.Next.Type
			p.Lex.SelectNext()
			if p.Lex.Next.Type != token.INT {
				panic("[Parser] Esperado um número inteiro, mas encontrado: " + p.Lex.Next.Type)
			}
			value := p.Lex.Next.Value.(int)

			switch operator {
				case token.PLUS:
					result += value
				case token.MINUS:
					result -= value
			}
		
			p.Lex.SelectNext()
	}
	return result
}