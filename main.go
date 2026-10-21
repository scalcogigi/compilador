package main

import (
	"compilador/code"
	"compilador/parser"
	"compilador/prepro"
	"compilador/symboltable"
	"compilador/node"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("uso: go run . <teste.txt>")
	}
	entrada := os.Args[1]

	conteudo, err := os.ReadFile(entrada)
	if err != nil {
		panic(err)
	}

	codigo := prepro.Filter(string(conteudo) + "\n")
	ast := parser.Run(codigo)

	node.Check(ast, symboltable.New())
	// gera instruções usando uma tabela nova
	code.Reset()
	st := symboltable.New()
	ast.Generate(st)
	code.SetStackSize(len(st.Table) * 4) // as variáveis ocupam 4 bytes inicialmente

	extensao := filepath.Ext(entrada)
	saida := strings.TrimSuffix(entrada, extensao) + ".asm"

	if err := code.Dump(saida); err != nil {
		panic(err)
	}
}

// processa os elementos, retorna o resultado e printa
