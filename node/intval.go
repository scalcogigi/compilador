package node

import "compilador/symboltable"

type IntVal struct {
	Value int
}

func (i IntVal) Evaluate(st symboltable.SymbolTable) int {
	return i.Value
}