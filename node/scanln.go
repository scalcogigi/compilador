package node

import (
	"compilador/symboltable"
	"fmt"
)

type ScanLn struct{}

func (s ScanLn) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	var value int

	if _, err := fmt.Scanln(&value); err != nil {
		panic("[Semantic] Erro ao ler inteiro: " + err.Error())
	}

	return symboltable.Variable{
		Value: value,
		Type:  "int",
	}
}