package node

import "compilador/symboltable"

type IntVal struct {
	Value int
}

func (i IntVal) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{
		Value: i.Value,
		Type: "int",
	}
}
