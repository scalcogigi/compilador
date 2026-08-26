package token

const (
	INT   = "INT" // número inteiro
	PLUS  = "PLUS" // +
	MINUS = "MINUS" // -
	EOF   = "EOF" // end of file
	XOR = "XOR" // ^
)

type Token struct {
	Type  string
	Value interface{} // int ou string
}	

