package node

import "compilador/symboltable"

type Assignment struct {
	Children []Node
}

func (a Assignment) Evaluate(st symboltable.SymbolTable) int {
	if len(a.Children) != 2 {
		panic("[Assignment] Esperados dois filhos")
	}

	identifier, ok := a.Children[0].(Identifier)
	if !ok {
		panic("[Assignment] Primeiro filho deve ser um Identifier")
	}

	value := a.Children[1].Evaluate(st)
	st.Insert(identifier.Value, value)

	return value
}
