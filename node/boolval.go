package node

import (
	"compilador/symboltable"
	"compilador/code"
)
type BoolVal struct {
	Base
	Value bool
}

func (b BoolVal) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{
		Value: b.Value,
		Type:  "bool",
	}
}

func (b BoolVal) Generate(st symboltable.SymbolTable) {
	if b.Value {
		code.Append("  mov eax, 1")
	} else {
		code.Append("  mov eax, 0")
	}
}