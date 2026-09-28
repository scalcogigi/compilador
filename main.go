package main

import (
	"compilador/parser"
	"compilador/prepro"
	"compilador/symboltable"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("uso: go run . \"1+2\"")
	}
	teste := os.Args[1]

	conteudo, err := os.ReadFile(teste)
	if err != nil {
		panic(err)
	}

	codigo := string(conteudo) + "\n"
	codigo = prepro.Filter(codigo)

	result := parser.Run(codigo)

	st := symboltable.New()
	result.Evaluate(st)
}

// processa os elementos, retorna o resultado e printa
