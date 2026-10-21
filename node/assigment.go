package node

import (
	"compilador/code"
	"compilador/symboltable"
	"fmt"
)

type Assignment struct {
	Base
	Children []Node
}

func (a Assignment) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(a.Children) != 2 {
		panic("[Assignment] Esperados dois filhos")
	}

	identifier, ok := a.Children[0].(Identifier)
	if !ok {
		panic("[Assignment] Primeiro filho deve ser um Identifier")
	}

	value := a.Children[1].Evaluate(st)
	st.Set(identifier.Value, value)

	return value
}

func (a Assignment) Generate(st symboltable.SymbolTable) {
	if len(a.Children) != 2 {
		panic("[Semantic] Assignment espera dois filhos")
	}

	identifier, ok := a.Children[0].(Identifier)
	if !ok {
		panic("[Semantic] Primeiro filho de Assignment deve ser um Identifier")
	}

	variable := st.Get(identifier.Value) // busca variável e deslocamento na pilha

	if variable.Type == "string" {
		panic("[Semantic] Geração de código não suporta strings")
	}

	a.Children[1].Generate(st) // resultado em eax

	code.Append(fmt.Sprintf("  mov [ebp-%d], eax", variable.Shift))
}
