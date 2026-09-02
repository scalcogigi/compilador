package main

import (
	"fmt"
	"os"
	"compilador/parser"
)

func main() {
	if len(os.Args) != 2 {
	panic("uso: go run . \"1+2\"")
	}
	expr := os.Args[1]
	result := parser.Run(expr)
	fmt.Println(result)
}

// processa os elementos, retorna o resultado e printa