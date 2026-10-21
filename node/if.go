package node

import (
	"compilador/code"
	"compilador/symboltable"
	"fmt"
)

type IF struct {
	Base
	Children []Node
}

func (i IF) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(i.Children) != 2 && len(i.Children) != 3 {
		panic("[IF] Esperados 2 ou 3 filhos")
	}

	condition := i.Children[0].Evaluate(st)
	if condition.Type != "bool" {
		panic("[Semantic] Condição do if deve ser bool")
	}

	if condition.Value.(bool) {
		i.Children[1].Evaluate(st)
	} else if len(i.Children) == 3 {
		i.Children[2].Evaluate(st)
	}	

	return symboltable.Variable{}
}

func (i IF) Generate(st symboltable.SymbolTable) {
	if len(i.Children) != 2 && len(i.Children) != 3 {
		panic("[Semantic] IF espera dois ou três filhos")
	}

	id := i.ID
	falseLabel := fmt.Sprintf("if_false_%d", id)
	endLabel := fmt.Sprintf("if_end_%d", id)

	// Gera a condição.
	i.Children[0].Generate(st)
	code.Append("  cmp eax, 0")
	code.Append("  je " + falseLabel)

	// Gera o bloco verdadeiro.
	i.Children[1].Generate(st)

	if len(i.Children) == 3 {
		// Impede que o bloco verdadeiro continue pelo else.
		code.Append("  jmp " + endLabel)
		code.Append(falseLabel + ":")

		i.Children[2].Generate(st)

		code.Append(endLabel + ":")
	} else {
		code.Append(falseLabel + ":")
	}
}