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
		
	case token.FACT:
		if value < 0 {
			panic("[Semantic] fatorial de número negativo")
		}

		result := 1 // !0 !1

		for i := 2; i <= value; i++ {
			result *= i
		}

		return result
		default:
			panic("[Semantic] operador unário inválido: " + u.Value)
		}
}