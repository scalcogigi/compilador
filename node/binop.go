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

func (b BinOp) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(b.Children) != 2 {
		panic("[Semantic] BinOp espera dois filhos")
	}

	left := b.Children[0].Evaluate(st)
	right := b.Children[1].Evaluate(st)

	switch b.Value {
	case token.PLUS:
		if left.Type == "string" && right.Type == "string" {
			return symboltable.Variable{
				Value: left.Value.(string) + right.Value.(string),
				Type:  "string",
			}
		}

		if left.Type == "int" && right.Type == "int" {
			return symboltable.Variable{
				Value: left.Value.(int) + right.Value.(int),
				Type:  "int",
			}
		}

		panic("[Semantic] Tipos incompatíveis para +")

	case token.MINUS, token.MULT, token.DIV, token.POWER:
		if left.Type != "int" || right.Type != "int" {
			panic("[Semantic] Operador aritmético exige dois int")
		}

		a := left.Value.(int)
		c := right.Value.(int)

		switch b.Value {
		case token.MINUS:
			return symboltable.Variable{Value: a - c, Type: "int"}

		case token.MULT:
			return symboltable.Variable{Value: a * c, Type: "int"}

		case token.DIV:
			if c == 0 {
				panic("[Semantic] Divisão por zero")
			}
			return symboltable.Variable{Value: a / c, Type: "int"}

		case token.POWER:
			result := int(math.Pow(float64(a), float64(c)))
			return symboltable.Variable{Value: result, Type: "int"}
		}
	

	case token.AND, token.OR:
		if left.Type != "bool" || right.Type != "bool" {
			panic("[Semantic] Operador lógico exige dois bool")
		}

		a := left.Value.(bool)
		c := right.Value.(bool)

		if b.Value == token.AND {
			return symboltable.Variable{Value: a && c, Type: "bool"}
		}
		return symboltable.Variable{Value: a || c, Type: "bool"}


	case token.EQ:
		if left.Type != right.Type {
			panic("[Semantic] Operador de igualdade exige dois valores do mesmo tipo")
		}
		return symboltable.Variable{
			Value: left.Value == right.Value,
			Type: "bool",
		}

	case token.LT, token.GT:
		if left.Type != "int" || right.Type != "int" {
			panic("[Semantic] Operador de comparação exige dois int")
		}

		a := left.Value.(int)
		c := right.Value.(int)

		if b.Value == token.LT {
			return symboltable.Variable{Value: a < c, Type: "bool"}
		}

		return symboltable.Variable{Value: a > c, Type: "bool"}

		}
		panic("[Semantic] operador binário inválido: " + b.Value)
}
