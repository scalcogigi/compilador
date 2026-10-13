package node

import "compilador/symboltable"

type StringVal struct {
	Value string
}

func (s StringVal) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{
		Value: s.Value,
		Type:  "string",
	}
}