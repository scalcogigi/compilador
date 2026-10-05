package node

import "compilador/symboltable"

type IF struct {
	Children []Node
}

func (i IF) Evaluate(st symboltable.SymbolTable) int {
	if len(i.Children) != 1 {
		panic("[IF] Esperado um filho")
	}

	condition := i.Children[0].Evaluate(st)
	if condition != 0 {
	}
	return 0
}

