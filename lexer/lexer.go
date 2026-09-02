package lexer

import (
	"compilador/token"
	"strconv"
)

// gera tokens - definindo variáveis e funções
type Lexer struct {
	Source   string
	Position int
	Next     token.Token
}

func (l *Lexer) SelectNext() { // seleciona o próximo token e remove espaços

	for l.Position < len(l.Source) && l.Source[l.Position] == ' ' {
		l.Position++
	}

	// se vazio
	if l.Position >= len(l.Source) {
		l.Next = token.Token{
			Type:  token.EOF,
			Value: "",
		}
		return
	}

	// se soma
	if l.Source[l.Position] == '+' {
		l.Next = token.Token{
			Type:  token.PLUS,
			Value: "+",
		}
		l.Position++
		return
	}

	// se subtrai
	if l.Source[l.Position] == '-' {
		l.Next = token.Token{
			Type:  token.MINUS,
			Value: "-",
		}
		l.Position++
		return
	}

	// se multiplica
	if l.Source[l.Position] == '*' {
		l.Next = token.Token{
			Type:  token.MULT,
			Value: "*",
		}
		l.Position++
		return
	}

	// se divide
	if l.Source[l.Position] == '/' {
		l.Next = token.Token{
			Type:  token.DIV,
			Value: "/",
		}
		l.Position++
		return
	}

	// se abre parêntese
	if l.Source[l.Position] == '(' {
		l.Next = token.Token{
			Type:  token.OPEN_PAREN,
			Value: "(",
		}
		l.Position++
		return
	}

	// se fecha parêntese
	if l.Source[l.Position] == ')' {
		l.Next = token.Token{
			Type:  token.CLOSE_PAREN,
			Value: ")",
		}
		l.Position++
		return
	}

	// se número, não deve ter ifs
	if l.Source[l.Position] >= '0' && l.Source[l.Position] <= '9' {
		number := "" // inicializa vazia para armanzenar o número

		for ; l.Position < len(l.Source) && l.Source[l.Position] >= '0' && l.Source[l.Position] <= '9'; l.Position++ {
			number += string(l.Source[l.Position])
		}
		numberInt, err := strconv.Atoi(number) // ASCII to int
		if err != nil {
			panic("[Lexer] Type must be a pointer, channel, func, interface, map, or slice type!")
		}

		l.Next = token.Token{
			Type:  token.INT,
			Value: numberInt,
		}

		return
	}

	// se inválido
	panic("[Lexer] Token inválido: " + string(l.Source[l.Position]))
}

// construtor
func New(source string) *Lexer {
	return &Lexer{ // ponteiro?
		Source:   source,
		Position: 0,
	}
}
