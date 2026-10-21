package node

import (
	"compilador/code"
	"compilador/symboltable"
	"fmt"
)

type VarDec struct {
	Base
	Value    string // tipo declarado: "int", "string" ou "bool"
	Children []Node // Identifier e, opcionalmente, expressão inicial
}

func (v VarDec) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(v.Children) != 1 && len(v.Children) != 2 {
		panic("[Semantic] Declaração deve ter identificador e, opcionalmente, valor inicial")
	}

	identifier, ok := v.Children[0].(Identifier)
	if !ok {
		panic("[Semantic] Primeiro filho deve ser um Identifier")
	}

	st.CreateVariable(identifier.Value, v.Value)

	if len(v.Children) == 2 {
		initialValue := v.Children[1].Evaluate(st)
		st.Set(identifier.Value, initialValue)
	}

	return symboltable.Variable{}
}

func (v VarDec) Generate(st symboltable.SymbolTable) {
	if len(v.Children) != 1 && len(v.Children) != 2 {
		panic("[Semantic] Declaração deve ter identificador e, opcionalmente, valor inicial")
	}

	identifier, ok := v.Children[0].(Identifier)
	if !ok {
		panic("[Semantic] Primeiro filho deve ser um Identifier")
	}

	if v.Value == "string" {
		panic("[Semantic] Geração de código não suporta strings")
	}

	st.CreateVariable(identifier.Value, v.Value)
	variable := st.Get(identifier.Value)

	code.Append(fmt.Sprintf("  mov dword [ebp-%d], 0", variable.Shift))

	if len(v.Children) == 2 {
		v.Children[1].Generate(st)
		code.Append(fmt.Sprintf("  mov [ebp-%d], eax", variable.Shift))
	}
}