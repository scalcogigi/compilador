package main

import (
	"fmt"
	"os"
	
	"compilador/parser"
	"compilador/symboltable"
	"compilador/prepro"
)

func main() {
	if len(os.Args) != 2 {
		panic("uso: go run . \"1+2\"")
	}
	expr := os.Args[1]
	result := parser.Run(expr)
	expr = prepro.Filter(expr)
	st := symboltable.New()
	fmt.Println(result.Evaluate(st))
}

// processa os elementos, retorna o resultado e printa