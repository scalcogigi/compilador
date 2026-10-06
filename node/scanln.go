package node

import (
	"compilador/symboltable"
	"fmt"
)

type ScanLn struct{}

func (s ScanLn) Evaluate(st symboltable.SymbolTable) int {
	var value int

	_, err := fmt.Scanln(&value)
	if err != nil {
		panic("[Semantic] Erro ao ler inteiro: " + err.Error())
	}

	return value
}