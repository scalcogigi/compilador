package node

// chamada após a construção da AST
func AssignIDs(root Node) Node {
	switch n := root.(type) {
	case Block:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case Assignment:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case VarDec:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case BinOp:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case UnOp:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case IF:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case While:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case Println:
		n.ID = NewID()
		n.Children = assignChildrenIDs(n.Children)
		return n

	case Identifier:
		n.ID = NewID()
		return n

	case IntVal:
		n.ID = NewID()
		return n

	case BoolVal:
		n.ID = NewID()
		return n

	case StringVal:
		n.ID = NewID()
		return n

	case ScanLn:
		n.ID = NewID()
		return n

	case NoOp:
		n.ID = NewID()
		return n

	case Empty:
		n.ID = NewID()
		return n

	default:
		panic("[AST] Tipo de nó sem atribuição de ID")
	}
}

func assignChildrenIDs(children []Node) []Node {
	result := make([]Node, len(children))

	for index, child := range children {
		result[index] = AssignIDs(child)
	}

	return result
}