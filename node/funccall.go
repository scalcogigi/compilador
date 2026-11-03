package node

import "compilador/symboltable"

type FUNCCALL struct {
	Value string
	Children []Node
}

func (i FUNCCALL) Evaluate(st symboltable.SymbolTable) symboltable.Variable {
	functionSymbol, exists := st.Find(i.Value)
	if !exists {
		panic("[Semantic] Função não definida: " + i.Value)
	}
	if !functionSymbol.IsFunction {
		panic("[Semantic] O identificador não é uma função: " + i.Value)
	}

	function, ok := functionSymbol.Value.(FUNCDEC)
	if !ok {
		panic("[Semantic] Referência inválida para a função: " + i.Value)
	}

	if len(function.Children) < 2 {
		panic("[Semantic] Declaração de função incompleta: " + i.Value)
	}

	parameterCount := len(function.Children) - 2
	if len(i.Children) != parameterCount {
		panic("[Semantic] Número incorreto de argumentos para: " + i.Value)
	}

	arguments := make([]symboltable.Variable, len(i.Children))
	for index, argument := range i.Children {
		arguments[index] = argument.Evaluate(st)
	}

	local := symboltable.New(&st)

	for index := 0; index < parameterCount; index++ {
		parameter, ok := function.Children[index+1].(VarDec)
		if !ok || len(parameter.Children) != 1 {
			panic("[Semantic] Parâmetro inválido na função: " + i.Value)
		}

		identifier, ok := parameter.Children[0].(Identifier)
		if !ok {
			panic("[Semantic] Nome de parâmetro inválido na função: " + i.Value)
		}

		if arguments[index].Type != parameter.Value {
			panic("[Semantic] Tipo incorreto para o parâmetro: " + identifier.Value)
		}

		local.CreateVariable(identifier.Value, parameter.Value)
		local.Set(identifier.Value, arguments[index])
	}

	body, ok := function.Children[len(function.Children)-1].(Block)
	if !ok {
		panic("[Semantic] Corpo inválido na função: " + i.Value)
	}

	result := body.Evaluate(local)

	if function.Value == "" {
		if result.IsReturn {
			panic("[Semantic] Função sem tipo de retorno não pode usar return: " + i.Value)
		}
		return symboltable.Variable{}
	}

	if !result.IsReturn {
		panic("[Semantic] Função terminou sem retornar um valor: " + i.Value)
	}
	if result.Type != function.Value {
		panic("[Semantic] Tipo de retorno incompatível na função: " + i.Value)
	}

	result.IsReturn = false
	return result
}