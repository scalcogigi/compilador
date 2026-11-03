package symboltable

type Variable struct {
	Value interface{}
	Type string
	IsReturn bool
	IsFunction bool
}
