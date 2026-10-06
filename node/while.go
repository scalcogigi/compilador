package node

import "compilador/symboltable"

type While struct {
	Children []Node
}

func (w While) Evaluate(st symboltable.SymbolTable) int {
	if len(w.Children) != 2 {
		panic("[Semantic] While espera dois filhos")
	}

	for w.Children[0].Evaluate(st) != 0 {
		w.Children[1].Evaluate(st)
	}

	return 0
}