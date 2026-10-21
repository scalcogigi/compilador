package node

import (
	"compilador/symboltable"
	"compilador/token"
)

// verifica tipos - verifica sem executar
func Check(root Node, st symboltable.SymbolTable) string {
	switch n := root.(type) {
	case IntVal:
		return "int"

	case BoolVal:
		return "bool"

	case StringVal:
		panic("[Semantic] Geração de código não suporta strings")

	case Identifier:
		variable := st.Get(n.Value)

		if variable.Type == "string" {
			panic("[Semantic] Geração de código não suporta strings")
		}

		return variable.Type

	case ScanLn:
		return "int"

	case Block:
		for _, child := range n.Children {
			Check(child, st)
		}
		return ""

	case NoOp:
		return ""

	case Empty:
		return ""

	case VarDec:
		if len(n.Children) != 1 && len(n.Children) != 2 {
			panic("[Semantic] VarDec espera um ou dois filhos")
		}

		identifier, ok := n.Children[0].(Identifier)
		if !ok {
			panic("[Semantic] Primeiro filho de VarDec deve ser Identifier")
		}

		if n.Value != "int" && n.Value != "bool" {
			panic("[Semantic] Tipo não suportado na geração: " + n.Value)
		}

		st.CreateVariable(identifier.Value, n.Value)

		if len(n.Children) == 2 {
			initialType := Check(n.Children[1], st)

			if initialType != n.Value {
				panic("[Semantic] Tipo incompatível na inicialização de: " + identifier.Value)
			}
		}

		return ""

	case Assignment:
		if len(n.Children) != 2 {
			panic("[Semantic] Assignment espera dois filhos")
		}

		identifier, ok := n.Children[0].(Identifier)
		if !ok {
			panic("[Semantic] Primeiro filho de Assignment deve ser Identifier")
		}

		declared := st.Get(identifier.Value)
		valueType := Check(n.Children[1], st)

		if declared.Type != valueType {
			panic("[Semantic] Tipo incompatível na atribuição para: " + identifier.Value)
		}

		return ""

	case Println:
		if len(n.Children) != 1 {
			panic("[Semantic] Println espera um filho")
		}

		valueType := Check(n.Children[0], st)
		if valueType != "int" && valueType != "bool" {
			panic("[Semantic] Println exige uma expressão int ou bool")
		}

		return ""

	case IF:
		if len(n.Children) != 2 && len(n.Children) != 3 {
			panic("[Semantic] IF espera dois ou três filhos")
		}

		if Check(n.Children[0], st) != "bool" {
			panic("[Semantic] Condição do if deve ser bool")
		}

		Check(n.Children[1], st)

		if len(n.Children) == 3 {
			Check(n.Children[2], st)
		}

		return ""

	case While:
		if len(n.Children) != 2 {
			panic("[Semantic] While espera dois filhos")
		}

		if Check(n.Children[0], st) != "bool" {
			panic("[Semantic] Condição do laço deve ser bool")
		}

		Check(n.Children[1], st)
		return ""

	case UnOp:
		if len(n.Children) != 1 {
			panic("[Semantic] UnOp espera um filho")
		}

		valueType := Check(n.Children[0], st)

		switch n.Value {
		case token.PLUS, token.MINUS:
			if valueType != "int" {
				panic("[Semantic] Operador unário aritmético exige int")
			}
			return "int"

		case token.NOT:
			if valueType != "bool" {
				panic("[Semantic] Operador ! exige bool")
			}
			return "bool"

		default:
			panic("[Semantic] Operador unário inválido: " + n.Value)
		}

	case BinOp:
		if len(n.Children) != 2 {
			panic("[Semantic] BinOp espera dois filhos")
		}

		leftType := Check(n.Children[0], st)
		rightType := Check(n.Children[1], st)

		switch n.Value {
		case token.PLUS, token.MINUS, token.MULT, token.DIV:
			if leftType != "int" || rightType != "int" {
				panic("[Semantic] Operador aritmético exige dois int")
			}
			return "int"

		case token.POWER:
			panic("[Semantic] Geração de potência ainda não implementada")

		case token.AND, token.OR:
			if leftType != "bool" || rightType != "bool" {
				panic("[Semantic] Operador lógico exige dois bool")
			}
			return "bool"

		case token.EQ:
			if leftType != rightType ||
				(leftType != "int" && leftType != "bool") {
				panic("[Semantic] Igualdade exige dois valores do mesmo tipo")
			}
			return "bool"

		case token.LT, token.GT:
			if leftType != "int" || rightType != "int" {
				panic("[Semantic] Comparação exige dois int")
			}
			return "bool"

		default:
			panic("[Semantic] Operador binário inválido: " + n.Value)
		}

	default:
		panic("[Semantic] Tipo de nó sem verificação implementada")
	}
}