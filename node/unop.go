package node

import (
	"compilador/symboltable"
	"compilador/token"
	"compilador/code"
)

type UnOp struct {
	Base
	Value    string
	Children []Node
}

func (u UnOp) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(u.Children) != 1 {
		panic("[UnOp] Esperado um filho")
	}

	value := u.Children[0].Evaluate(st) // fica no primeiro filho

	switch u.Value {

	case token.PLUS:
		if value.Type != "int" {
			panic("[Semantic] operador unário '+' exige int")
		}
		return value

	case token.MINUS:
		if value.Type != "int" {
			panic("[Semantic] operador unário '-' exige int")
		}
		return symboltable.Variable{
			Value: -value.Value.(int),
			Type:  "int",
		}

	case token.NOT:
		if value.Type != "bool" {
			panic("[Semantic] Operador ! exige bool")
		}
		return symboltable.Variable{
			Value: !value.Value.(bool),
			Type:  "bool",
		}


	default:
		panic("[Semantic] operador unário inválido: " + u.Value)
	}
}

func (u UnOp) Generate(st symboltable.SymbolTable) {
	if len(u.Children) != 1 {
		panic("[Semantic] UnOp espera um filho")
	}

	u.Children[0].Generate(st)

	switch u.Value {
	case token.PLUS: // valor no EAX por definição

	case token.MINUS:
		code.Append("  neg eax")

	case token.NOT:
		code.Append("  cmp eax, 0")
		code.Append("  sete al")
		code.Append("  movzx eax, al")

	default:
		panic("[Semantic] Operador unário inválido: " + u.Value)
	}
}