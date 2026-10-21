package node

import "compilador/symboltable"

type Empty struct{
	Base
}

func (e Empty) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	return symboltable.Variable{}
}

func (e Empty) Generate(st symboltable.SymbolTable) { // não emite assembly
}