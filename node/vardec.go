package node

import "compilador/symboltable"

type VarDec struct {
	Value    string // tipo declarado: "int", "string" ou "bool"
	Children []Node // Identifier e, opcionalmente, expressão inicial
}

func (v VarDec) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(v.Children) != 1 && len(v.Children) != 2 {
		panic("[VarDec] Declaração deve ter identificador e, opcionalmente, valor inicial")
	}

	identifier, ok := v.Children[0].(Identifier)
	if !ok {
		panic("[VarDec] Primeiro filho deve ser um Identifier")
	}

	st.CreateVariable(identifier.Value, v.Value)

	if len(v.Children) == 2 {
		initialValue := v.Children[1].Evaluate(st)
		st.Set(identifier.Value, initialValue)
	}

	return symboltable.Variable{}
}