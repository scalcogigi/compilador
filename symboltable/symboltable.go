package symboltable

type SymbolTable struct {
	Table map[string]Variable
}

// método para criar nova tabela
func New() SymbolTable {
	return SymbolTable{
		Table: make(map[string]Variable),
	}
}

func (st SymbolTable) CreateVariable(name string, variableType string) {
	if _, exists := st.Table[name]; exists {
		panic("[Semantic] Variável já declarada: " + name)
	}

	var defaultValue interface{}

	switch variableType {
	case "int":
		defaultValue = 0
	case "string":
		defaultValue = ""
	case "bool":
		defaultValue = false
	default:
		panic("[Semantic] Tipo inválido: " + variableType)
	}

	st.Table[name] = Variable{
		Value: defaultValue,
		Type:  variableType,
		Shift: (len(st.Table) + 1) * 4,
		// EBP é um registrador usado como referência para acessar as variáveis na pilha
	}
}

// Altera uma variável que já foi declarada.
func (st SymbolTable) Set(name string, value Variable) {
	declared, exists := st.Table[name]
	if !exists {
		panic("[Semantic] Variável não definida: " + name)
	}

	if declared.Type != value.Type {
		panic("[Semantic] Tipo incompatível na atribuição para: " + name)
	}

	declared.Value = value.Value
	st.Table[name] = declared

}

// método para buscar
func (st *SymbolTable) Find(name string) (Variable, bool) {
	variable, exists := st.Table[name]
	return variable, exists
}

func (st SymbolTable) Get(name string) Variable {
	variable, exists := st.Find(name)

	if !exists {
		panic("[Semantic] Variável não definida: " + name)
	}

	return variable
}
