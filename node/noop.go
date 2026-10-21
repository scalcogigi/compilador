package node

import "compilador/symboltable"

type NoOp struct{
	Base
}

func (n NoOp) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{}
}

func (n NoOp) Generate(st symboltable.SymbolTable) { // no operation literalmente
}