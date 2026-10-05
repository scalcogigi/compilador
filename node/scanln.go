package node

import "compilador/symboltable"

type ScanLn struct {
	Children []Node
}

func (s ScanLn) Evaluate(st symboltable.SymbolTable) int {
	if len(s.Children) != 1 {
		panic("[ScanLn] Esperado um filho")
	}
	return 0
}