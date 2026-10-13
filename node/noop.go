package node

import "compilador/symboltable"

type NoOp struct{}

func (n NoOp) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{}
}
