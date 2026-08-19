package parser
import ("strconv")

func Run(code string) int {
	
	code = removeSpaces(code)
	result, pos := readNumber(code, 0)

	for pos < len(code) {
		operator := code[pos]
		pos++
		if pos >= len(code) {
    		panic("esperava um número")
		}		
		nextNumber, newPos := readNumber(code, pos)
		if operator == '+' {
			result += nextNumber
		} else if operator == '-' {
			result -= nextNumber
		} else {
			panic("Operador inválido: " + string(operator))
		}
		pos = newPos
	}
	return result
} 

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// pré-processamento 
func removeSpaces(code string) (string) {
	result := ""
    for i := 0; i < len(code); i++ {
        c := code[i]
        if c == ' ' { // remove espaços
			if i > 0 &&
				i < len(code)-1 &&
				isDigit(code[i-1]) &&
				isDigit(code[i+1]) {

        panic("esperava um operador")
    }

    continue
}

	result += string(c)
	}
    return result
}

// considera apenas números
func readNumber(code string, start int) (int, int) {
	number := ""
	i := start
	for ; i < len(code); i++ {
		if code[i] >= '0' && code[i] <= '9'{ // como é um byte, faz sentido verificar todos os dígitos
			number += string(code[i])
		} else {
			break
		}
	}

	numberInt, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return numberInt, i
}