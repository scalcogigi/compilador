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

	if l.Position >= len(l.Source) {
		l.Next = token.Token{
			Type:  token.EOF,
			Value: "",
		}
		return
	}

	if l.Source[l.Position] == '+' {
		l.Next = token.Token{
			Type:  token.PLUS,
			Value: "+",
		}
		l.Position++
		return
	}

	if l.Source[l.Position] == '-' {
		l.Next = token.Token{
			Type:  token.MINUS,
			Value: "-",
		}
		l.Position++
		return
	}

	if l.Source[l.Position] == '^' {
		l.Next = token.Token{
			Type:  token.XOR,
			Value: "^",
		}
		l.Position++
		return
	}

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

	panic("[Lexer] Token inválido: " + string(l.Source[l.Position]))
}

// construtor
func New(source string) *Lexer {
	return &Lexer{ // ponteiro?
		Source:   source,
		Position: 0,
	}
}
