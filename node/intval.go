package node

import (
	"compilador/code"
	"compilador/symboltable"
	"fmt"
)

type IntVal struct {
	Base
	Value int
}

func (i IntVal) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{
		Value: i.Value,
		Type: "int",
	}
}

func (i IntVal) Generate(st symboltable.SymbolTable) {
	code.Append(fmt.Sprintf("  mov eax, %d", i.Value))
}