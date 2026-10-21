package code


import (
  "fmt"
	"os"
	"strings"
)

var instructions []string
var stackSize int // bytes das variáveis

func Reset() { // limpeza
	instructions = nil
  stackSize = 0
}

func Append(instruction string) { // adiciona novo assembly
	instructions = append(instructions, instruction)
}

func SetStackSize(size int) {
	stackSize = size
}

func Dump(filename string) error { // salva o programa
	header := `section .data
  format_out: db "%d", 10, 0
  format_in: db "%d", 0
  scan_int: dd 0

section .text
  extern printf
  extern scanf
  extern exit
  global _start

_start:
  push ebp
  mov ebp, esp

`

	if stackSize > 0 {
		header += fmt.Sprintf("  sub esp, %d\n\n", stackSize)
	}

	footer := `
  mov esp, ebp
  pop ebp

  push dword 0
  call exit
`

	content := header + strings.Join(instructions, "\n") + "\n" + footer

	return os.WriteFile(filename, []byte(content), 0644)
}