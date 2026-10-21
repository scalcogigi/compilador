package node

import (
	"compilador/symboltable"
	"compilador/code"
	"fmt"
)

type While struct {
	Base
	Children []Node
}

func (w While) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(w.Children) != 2 {
		panic("[Semantic] While espera dois filhos")
	}

	for {
		condition := w.Children[0].Evaluate(st)
		if condition.Type != "bool" {
			panic("[Semantic] Condição do laço deve ser bool")
		}

		if !condition.Value.(bool) {
			break
		}

		w.Children[1].Evaluate(st)
	}

	return symboltable.Variable{}
}

func (w While) Generate(st symboltable.SymbolTable) {
	if len(w.Children) != 2 {
		panic("[Semantic] While espera dois filhos")
	}

	id := w.ID
	loopLabel := fmt.Sprintf("loop_%d", id)
	exitLabel := fmt.Sprintf("exit_%d", id)

	code.Append(loopLabel + ":")

	w.Children[0].Generate(st)
	code.Append("  cmp eax, 0")
	code.Append("  je " + exitLabel)

	w.Children[1].Generate(st)

	code.Append("  jmp " + loopLabel)
	code.Append(exitLabel + ":")
}