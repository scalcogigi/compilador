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

// método para inserir
func (st *SymbolTable) Insert(name string, value int) {
	st.Table[name] = Variable{Value: value}
}

// método para buscar
func (st *SymbolTable) Find(name string) (Variable, bool) {
	variable, exists := st.Table[name]
	return variable, exists
}

func (st SymbolTable) Get(name string) int {
	variable, exists := st.Find(name)

	if !exists {
		panic("[SymbolTable] Variável não definida: " + name)
	}

	return variable.Value
}
