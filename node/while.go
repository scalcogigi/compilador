package node

import "compilador/symboltable"

type While struct {
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