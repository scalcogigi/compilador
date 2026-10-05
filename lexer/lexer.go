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

	for l.Position < len(l.Source) && (l.Source[l.Position] == ' ' ||
		l.Source[l.Position] == '\t' || l.Source[l.Position] == '\r') {
		l.Position++
	}

	// Quebra de linha indica o fim de uma instrução.
	if l.Position < len(l.Source) && l.Source[l.Position] == '\n' {
		l.Next = token.Token{
			Type:  token.END,
			Value: "\n",
		}
		l.Position++

		return
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

	// se AND
	if l.Position+1 < len(l.Source) &&
		l.Source[l.Position:l.Position+2] == "&&" {

		l.Next = token.Token{
			Type:  token.AND,
			Value: "&&",
		}
		l.Position += 2

		return
	}

	// se OR
	if l.Position+1 < len(l.Source) &&
		l.Source[l.Position:l.Position+2] == "||" {

		l.Next = token.Token{
			Type:  token.OR,
			Value: "||",
		}
		l.Position += 2

		return
	}

	// se potencia
	if l.Position+1 < len(l.Source) &&
		l.Source[l.Position:l.Position+2] == "**" {

		l.Next = token.Token{
			Type:  token.POWER,
			Value: "**",
		}
		l.Position += 2

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

	// se EQ
	if l.Position+1 < len(l.Source) &&
		l.Source[l.Position:l.Position+2] == "==" {

		l.Next = token.Token{
			Type:  token.EQ,
			Value: "==",
		}
		l.Position += 2

		return
	}

	// se atribui
	if l.Source[l.Position] == '=' {
		l.Next = token.Token{
			Type:  token.ASSIGN,
			Value: "=",
		}
		l.Position++

		return
	}

	// maior que
	if l.Source[l.Position] == '>' {
		l.Next = token.Token{
			Type:  token.GT,
			Value: ">",
		}
		l.Position++

		return
	}

	// menor que
	if l.Source[l.Position] == '<' {
		l.Next = token.Token{
			Type:  token.LT,
			Value: "<",
		}
		l.Position++

		return
	}

	// se not
	if l.Source[l.Position] == '!' {
		l.Next = token.Token{
			Type:  token.NOT,
			Value: "!",
		}
		l.Position++

		return
	}

	// se abre chave
	if l.Source[l.Position] == '{' {
		l.Next = token.Token{
			Type:  token.OPEN_BRA,
			Value: "{",
		}
		l.Position++

		return
	}

	// se fecha chave
	if l.Source[l.Position] == '}' {
		l.Next = token.Token{
			Type:  token.CLOSE_BRA,
			Value: "}",
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

	// Se for identificador ou palavra reservada.
	if isLetter(l.Source[l.Position]) {
		identifier := ""

		// Continua enquanto houver letra, número ou "_".
		for l.Position < len(l.Source) &&
			(isLetter(l.Source[l.Position]) ||
				isDigit(l.Source[l.Position]) ||
				l.Source[l.Position] == '_') {

			identifier += string(l.Source[l.Position])
			l.Position++
		}

		switch identifier {
			case "Println":
				// token PRINT
				l.Next = token.Token{
				Type:  token.PRINT,
				Value: identifier,
			}
			case "if":
				// token IF
				l.Next = token.Token{
				Type:  token.IF,
				Value: identifier,
			}

			case "for":
				// token WHILE (ou o token definido no seu token.go)
				l.Next = token.Token{
				Type:  token.WHILE,
				Value: identifier,
			}

			case "else":
				// token ELSE
				l.Next = token.Token{
				Type:  token.ELSE,
				Value: identifier,
			}

			case "Scanln":
				// token READ
				l.Next = token.Token{
				Type:  token.READ,
				Value: identifier,
			}

			default:
				// token IDEN
				l.Next = token.Token{
				Type:  token.IDEN,
				Value: identifier,
			}
			return 
		}
	}

	// se número, não deve ter ifs
	if isDigit(l.Source[l.Position]) {
		number := ""

		for l.Position < len(l.Source) && isDigit(l.Source[l.Position]) {
			number += string(l.Source[l.Position])
			l.Position++
		}
		numberInt, err := strconv.Atoi(number) // ASCII to int
		if err != nil {
			panic("[Lexer] Invalid number!")
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

// funções auxiliares
func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
