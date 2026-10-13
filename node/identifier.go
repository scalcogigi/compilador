package node

import "compilador/symboltable"

type Identifier struct {
	Value string
}

func (i Identifier) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return st.Get(i.Value)
}
