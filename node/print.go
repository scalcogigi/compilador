package node

import (
	"compilador/code"
	"compilador/symboltable"
	"fmt"
)

type Println struct {
	Base
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

func (p Println) Generate(st symboltable.SymbolTable) {
	if len(p.Children) != 1 {
		panic("[Semantic] Println espera um filho")
	}

	p.Children[0].Generate(st) // gera expressão

	code.Append("  push eax")
	code.Append("  push format_out")
	code.Append("  call printf")

	code.Append("  add esp, 8") // remove os dois argumentos de 4 bytes
}