package node

import "compilador/symboltable"

type IF struct {
	Children []Node
}

func (i IF) Evaluate(st symboltable.SymbolTable) int {
	if len(i.Children) != 2 && len(i.Children) != 3 {
		panic("[IF] Esperados 2 ou 3 filhos")
	}

	condition := i.Children[0].Evaluate(st)

	if condition != 0 {
		i.Children[1].Evaluate(st)
	} else if len(i.Children) == 3 {
		i.Children[2].Evaluate(st)
	}

	return 0
}