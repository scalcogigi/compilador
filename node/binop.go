package node

import (
	"compilador/symboltable"
	"compilador/token"
	"math"
)

type BinOp struct {
	Value    string
	Children []Node
}

func (b BinOp) Evaluate(st symboltable.SymbolTable) int {
	leftValue := b.Children[0].Evaluate(st)
	rightValue := b.Children[1].Evaluate(st)

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
	case token.POWER:
		return int(math.Pow(
			float64(leftValue),
			float64(rightValue),
		))
	default:
		panic("[Semantic] operador binário inválido: " + b.Value)
	}
}
