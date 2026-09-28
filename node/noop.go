package node

import "compilador/symboltable"

type NoOp struct{}

func (n NoOp) Evaluate(st symboltable.SymbolTable) int {
	return 0
}
