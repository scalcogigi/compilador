package node

import "compilador/symboltable"

type Block struct {
	Children []Node
}

func (b Block) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	var result symboltable.Variable

	for _, child := range b.Children {
		if nestedBlock, ok := child.(Block); ok {
			local := symboltable.New(&st)
			result = nestedBlock.Evaluate(local)
		} else {
			result = child.Evaluate(st)
		}

		if result.IsReturn {
			return result
		}
	}

	return result
}