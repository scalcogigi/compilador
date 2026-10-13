package node

import (
	"compilador/symboltable"
	"fmt"
)

type Println struct {
	Children []Node
}

func (p Println) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	if len(p.Children) != 1 {
		panic("[Println] Esperado um filho")
	}
	value := p.Children[0].Evaluate(st)
	fmt.Println(value.Value)
	return value
}
