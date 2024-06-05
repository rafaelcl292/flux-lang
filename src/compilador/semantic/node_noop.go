package semantic

type NoOp struct{}

func (n NoOp) Eval(st *SymbolTable) symbol {
	return symbol{NONE, 0}
}
