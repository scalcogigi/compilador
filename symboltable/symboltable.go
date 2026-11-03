package symboltable

type SymbolTable struct {
	Table map[string]Variable
	Parent *SymbolTable
}

// método para criar nova tabela
func New(parent *SymbolTable) SymbolTable {
	return SymbolTable{
		Table: make(map[string]Variable),
		Parent: parent,
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
	}
}

// Altera uma variável que já foi declarada, busca recursiva
func (st SymbolTable) Set(name string, value Variable) {
	declared, exists := st.Table[name]

	if exists {
		if declared.IsFunction {
			panic("[Semantic] Não é possível atribuir à função: " + name)
		}

		if declared.Type != value.Type {
			panic("[Semantic] Tipo incompatível na atribuição para: " + name)
		}

		st.Table[name] = value
		return
	}

	if st.Parent != nil {
		st.Parent.Set(name, value)
		return
	}

	panic("[Semantic] Variável não definida: " + name)
}

// método para buscar
func (st *SymbolTable) Find(name string) (Variable, bool) {
	variable, exists := st.Table[name]
	if exists {
		return variable, true
	}

	if st.Parent != nil {
		return st.Parent.Find(name)
	}

	return Variable{}, false
}

func (st SymbolTable) Get(name string) Variable {
	variable, exists := st.Find(name)

	if !exists {
		panic("[Semantic] Variável não definida: " + name)
	}

	return variable
}
