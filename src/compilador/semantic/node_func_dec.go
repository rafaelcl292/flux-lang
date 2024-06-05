package semantic

type FuncDec struct {
	Ident     string
	Vars      []Assign
	FuncBlock Block
}

func (n FuncDec) Eval(st *SymbolTable) symbol {
	ft.set(n.Ident, n)
	return symbol{NONE, nil}
}
