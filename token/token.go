package token

const (
	INT         = "INT"         // número inteiro
	PLUS        = "PLUS"        // +
	MINUS       = "MINUS"       // -
	EOF         = "EOF"         // end of file
	MULT        = "MULT"        // *
	DIV         = "DIV"         // /
	OPEN_PAREN  = "OPEN_PAREN"  // (
	CLOSE_PAREN = "CLOSE_PAREN" // )
	POWER = "POWER" // **
)

type Token struct {
	Type  string
	Value interface{} // int ou string
}
