package node

import ("compilador/token")

type UnOp struct {
	Value string
	Children []Node
}

func (u UnOp) Evaluate() int {
	value := u.Children[0].Evaluate() // fica no primeiro filho

	switch u.Value {
	case token.PLUS:
		return value
	case token.MINUS:
		return -value
	default:
		panic("[Semantic] operador unário inválido: " + u.Value)
	}
}