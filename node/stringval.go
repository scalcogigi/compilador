package node

import "compilador/symboltable"

type StringVal struct {
	Base
	Value string
}

func (s StringVal) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{
		Value: s.Value,
		Type:  "string",
	}
}

func (s StringVal) Generate(st symboltable.SymbolTable) {
	panic("[Semantic] Geração de código não suporta strings")
}