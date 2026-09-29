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
	POWER       = "POWER"       // **
	FACT        = "FACT"        // !
	ASSIGN      = "ASSIGN"      // =
	END         = "END"         // \n
	PRINT       = "PRINT"       // print
	IDEN        = "IDEN"        // variável
	AND         = "AND"         // and
	OR          = "OR"          // or
	NOT         = "NOT"         // not
	EQ          = "EQ"          // ==
	LT          = "LT"          // <
	GT          = "GT"          // >
	IF          = "IF"          // if
	WHILE       = "WHILE"       // while
	ELSE        = "ELSE"        // else
	READ        = "READ"        // read
	OPEN_BRA    = "OPEN_BRA"    // {
	CLOSE_BRA   = "CLOSE_BRA"   // }
)

type Token struct {
	Type  string
	Value interface{} // int ou string
}

