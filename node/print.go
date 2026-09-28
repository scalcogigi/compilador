package node

import (
	"compilador/symboltable"
)

type Println struct {
	Children []Node
}

func (p Println) Evaluate(st symboltable.SymbolTable) int {
	if len(p.Children) != 1 {
		panic("[Println] Esperado um filho")
	}
	value := p.Children[0].Evaluate(st)
	println(value)
	return value
}
