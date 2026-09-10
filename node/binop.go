package node

import ("compilador/token")

type BinOp struct {
	Value string
	Children []Node
}

func (b BinOp) Evaluate() int {
	leftValue := b.Children[0].Evaluate()
	rightValue := b.Children[1].Evaluate()

	switch b.Value {
	case token.PLUS:
		return leftValue + rightValue
	case token.MINUS:
		return leftValue - rightValue
	case token.MULT:
		return leftValue * rightValue
	case token.DIV:
		if rightValue == 0 {
			panic("[Semantic] divisão por zero")
		}
		return leftValue / rightValue
	default:
		panic("[Semantic] operador binário inválido: " + b.Value)
	}
}
