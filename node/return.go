package node

import "compilador/symboltable"

type RETURN struct {
	Children []Node
}

func (r RETURN) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(r.Children) != 1 {
		panic("[Return] Deve haver exatamente uma expressão")
	}

	result := r.Children[0].Evaluate(st)
	result.IsReturn = true

	return result
}