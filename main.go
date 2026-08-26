package main

import (
	"fmt"
	"os"
	"compilador/parser"
)

func main() {
	expr := os.Args[1]
	result := parser.Run(expr)
	fmt.Println(result)
}

// processa os elementos, retorna o resultado e printa