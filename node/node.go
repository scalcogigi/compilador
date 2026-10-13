package node

import "compilador/symboltable"

type Node interface {
	Evaluate(st symboltable.SymbolTable) symboltable.Variable
}
