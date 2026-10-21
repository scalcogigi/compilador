package node

import (
	"compilador/code"
	"compilador/symboltable"
	"fmt"
)

type ScanLn struct{
	Base
}

func (s ScanLn) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	var value int

	if _, err := fmt.Scanln(&value); err != nil {
		panic("[Semantic] Erro ao ler inteiro: " + err.Error())
	}

	return symboltable.Variable{
		Value: value,
		Type:  "int",
	}
}

func (s ScanLn) Generate(st symboltable.SymbolTable) {
	// passa endereço e o formato
	code.Append("  push scan_int")
	code.Append("  push format_in")
	code.Append("  call scanf")

	code.Append("  add esp, 8") // remove dois argumentos da pilha

	code.Append("  mov eax, dword [scan_int]")
}