package node

import "compilador/symboltable"

type Empty struct{}

func (e Empty) Evaluate(st symboltable.SymbolTable) int {
	return 0
}
