package node

type Node interface {
    Evaluate(st *SymbolTable) int
}