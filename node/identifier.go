package node

import (
	"compilador/code"
	"compilador/symboltable"
	"fmt"
)

type Identifier struct {
	Base
	Value string
}

func (i Identifier) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return st.Get(i.Value)
}

func (i Identifier) Generate(st symboltable.SymbolTable) {
	variable := st.Get(i.Value)

	if variable.Type == "string" { // quer mudar o valor que está no registrador (binário)
		panic("[Semantic] Geração de código não suporta strings")
	}

	code.Append(fmt.Sprintf("  mov eax, [ebp-%d]", variable.Shift)) // o valor muda dependendo di tamanho do pulo
}