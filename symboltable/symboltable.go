package symboltable

type SymbolTable struct {
    Table map[string]Variable
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