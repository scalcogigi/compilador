package node

import "compilador/symboltable"

type Block struct {
	Children []Node
}

func (b Block) Evaluate(st symboltable.SymbolTable) int {
	result := 0

	for _, child := range b.Children {
		result = child.Evaluate(st)
	}

	return result
}
