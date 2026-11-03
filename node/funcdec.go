package node

import "compilador/symboltable"

type FUNCDEC struct {
	Value string
	Children []Node
}

func (i FUNCDEC) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(i.Children) < 2 {
		panic("[Semantic] Declaração de função incompleta")
	}

	identifier, ok := i.Children[0].(Identifier)
	if !ok {
		panic("[Semantic] O primeiro filho de FuncDec deve ser um Identifier")
	}

	if _, exists := st.Table[identifier.Value]; exists {
		panic("[Semantic] Nome já declarado: " + identifier.Value)
	}

	st.Table[identifier.Value] = symboltable.Variable{
		Value:      i,
		Type:       i.Value,
		IsFunction: true,
	}
	
	return symboltable.Variable{}
}