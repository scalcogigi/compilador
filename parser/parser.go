package parser
import "fmt"

func Run(code string) int {
	
	code = removeSpaces(code)
	number, _ := readNumber(code, 0)

    fmt.Println(number)
	return len(code)
}

// pré-processamento 
func removeSpaces(code string) (string) {
	result := ""
    for _, c := range code {
        if c != ' ' { // remove espaços
            result += string(c)
        }
    }
    return result
}

// considera apenas números
func readNumber(code string, start int) (string, int) {
	number := ""
	i := start
	for ; i < len(code); i++ {
		if code[i] >= '0' && code[i] <= '9'{ // como é um byte, faz sentido verificar todos os dígitos
			number += string(code[i])
		} else {
			break
		}
	}
	return number, i
}