package token

const (
	INT   = "INT" // número inteiro
	PLUS  = "PLUS" // +
	MINUS = "MINUS" // -
	EOF   = "EOF" // end of file
	MULT = "MULT" // *
	DIV = "DIV" // /
	LPAREN = "LPAREN" // (
	RPAREN = "RPAREN" // )
)

type Token struct {
	Type  string
	Value interface{} // int ou string
}	

