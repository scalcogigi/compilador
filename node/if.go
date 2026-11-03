package node

import "compilador/symboltable"

type IF struct {
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
		local := symboltable.New(&st)
		return i.Children[1].Evaluate(local)
	}

	if len(i.Children) == 3 {
		local := symboltable.New(&st)
		return i.Children[2].Evaluate(local)
	}

	return symboltable.Variable{}
}