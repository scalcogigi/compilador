package node

import "compilador/symboltable"

type Identifier struct {
	Value string
}

func (i Identifier) Evaluate(st symboltable.SymbolTable) int {
	return st.Get(i.Value)
}
