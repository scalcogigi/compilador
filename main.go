package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("len(os.Args) = %d\n", len(os.Args))
	fmt.Printf("%q\n", os.Args)
}