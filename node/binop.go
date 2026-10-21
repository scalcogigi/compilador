package node

import (
	"compilador/symboltable"
	"compilador/token"
	"compilador/code"
	"fmt"
	"math"
)

type BinOp struct {
	Base
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
		if left.Type == "int" && right.Type == "int" {
			return symboltable.Variable{
				Value: left.Value.(int) + right.Value.(int),
				Type:  "int",
			}
		}

		if left.Type == "string" || right.Type == "string" {
			return symboltable.Variable{
				Value: fmt.Sprint(left.Value) + fmt.Sprint(right.Value),
				Type:  "string",
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
			return symboltable.Variable{
				Value: int(math.Pow(float64(a), float64(c))),
				Type:  "int",
			}
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
			Type:  "bool",
		}

	case token.LT, token.GT:
		if left.Type == "int" && right.Type == "int" {
			a := left.Value.(int)
			c := right.Value.(int)
			if b.Value == token.LT {
				return symboltable.Variable{Value: a < c, Type: "bool"}
			}
			return symboltable.Variable{Value: a > c, Type: "bool"}
		}

		if left.Type == "string" && right.Type == "string" {
			a := left.Value.(string)
			c := right.Value.(string)
			if b.Value == token.LT {
				return symboltable.Variable{Value: a < c, Type: "bool"}
			}
			return symboltable.Variable{Value: a > c, Type: "bool"}
		}

		panic("[Semantic] Tipos incompatíveis para comparação")
	}

	panic("[Semantic] Operador binário inválido: " + b.Value)
}

func (b BinOp) Generate(st symboltable.SymbolTable) {
	if len(b.Children) != 2 {
		panic("[Semantic] BinOp espera dois filhos")
	}

	b.Children[0].Generate(st)
	code.Append("  push eax")

	b.Children[1].Generate(st)
	code.Append("  mov ecx, eax")
	code.Append("  pop eax")

	// eax e ecx
	switch b.Value {
	case token.PLUS:
		code.Append("  add eax, ecx")

	case token.MINUS:
		code.Append("  sub eax, ecx")

	case token.MULT:
		code.Append("  imul eax, ecx")

	case token.DIV:
		code.Append("  cdq")
		code.Append("  idiv ecx")

	case token.AND:
		code.Append("  cmp eax, 0")
		code.Append("  setne al")
		code.Append("  movzx eax, al")

		code.Append("  cmp ecx, 0")
		code.Append("  setne cl")
		code.Append("  movzx ecx, cl")

		code.Append("  and eax, ecx")

	case token.OR:
		code.Append("  or eax, ecx")
		code.Append("  setne al")
		code.Append("  movzx eax, al")

	case token.EQ:
		code.Append("  cmp eax, ecx")
		code.Append("  sete al")
		code.Append("  movzx eax, al")

	case token.LT:
		code.Append("  cmp eax, ecx")
		code.Append("  setl al")
		code.Append("  movzx eax, al")

	case token.GT:
		code.Append("  cmp eax, ecx")
		code.Append("  setg al")
		code.Append("  movzx eax, al")

	default:
		panic("[Semantic] Operador binário inválido: " + b.Value)
		}
	}