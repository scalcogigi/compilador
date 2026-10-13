package node

import "compilador/symboltable"

type BoolVal struct {
	Value bool
}

func (b BoolVal) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{
		Value: b.Value,
		Type:  "bool",
	}
}