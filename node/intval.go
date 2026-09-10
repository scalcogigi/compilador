package node

type IntVal struct {
	Value int
}

func (i IntVal) Evaluate() int {
	return i.Value
}