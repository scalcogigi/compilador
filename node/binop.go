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

	case token.AND:
		// Verdadeiro se ambos forem diferentes de zero.
		if leftValue != 0 && rightValue != 0 {
			return 1
		}
		return 0

	case token.OR:
		// Verdadeiro se pelo menos um for diferente de zero.
		if leftValue != 0 || rightValue != 0 {
			return 1
		}
		return 0

	case token.EQ:
		if leftValue == rightValue {
			return 1
		}
		return 0

	case token.LT:
		if leftValue < rightValue {
			return 1
		}
		return 0

	case token.GT:
		if leftValue > rightValue {
			return 1
		}
		return 0

	default:
		panic("[Semantic] operador binário inválido: " + b.Value)
	}
}
