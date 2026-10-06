package node

import (
	"compilador/symboltable"
	"compilador/token"
)

type UnOp struct {
	Value    string
	Children []Node
}

func (u UnOp) Evaluate(st symboltable.SymbolTable) int {
	value := u.Children[0].Evaluate(st) // fica no primeiro filho

	switch u.Value {

	case token.PLUS:
		return value

	case token.MINUS:
		return -value

	case token.NOT:
		if value == 0 {
			return 1
		}
		return 0

	default:
		panic("[Semantic] operador unário inválido: " + u.Value)
	}
}
