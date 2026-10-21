package node

import "compilador/symboltable"

type Block struct {
	Base
	Children []Node
}

func (b Block) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	var result symboltable.Variable

	for _, child := range b.Children {
		result = child.Evaluate(st)
	}

	return result
}

func (b Block) Generate(st symboltable.SymbolTable) { // percorre os todos os filhos, assim como o Evaluate
	for _, child := range b.Children {
		child.Generate(st)
	}
}