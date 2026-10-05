package node

import "compilador/symboltable"

type While struct {
	Children []Node
}

func (a While) Evaluate(st symboltable.SymbolTable) int {
	if len(a.Children) != 2 {
		panic("[While] Esperados dois filhos")
	}

	identifier, ok := a.Children[0].(Identifier)
	if !ok {
		panic("[While] Primeiro filho deve ser um Identifier")
	}

	value := a.Children[1].Evaluate(st)
	st.Insert(identifier.Value, value)

	return value
}
