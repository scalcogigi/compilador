package node

import "compilador/symboltable"

type Node interface {
	Evaluate(st symboltable.SymbolTable) symboltable.Variable
	Generate(st symboltable.SymbolTable)
}

type Base struct {
	ID int
}

var nextID int // contador compartilhado

func NewID() int {
	nextID++
	return nextID
}
