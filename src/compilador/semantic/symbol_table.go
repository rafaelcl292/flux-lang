package semantic

type SymbolTable map[string]symbol
type funcTable map[string]FuncDec

var ft = make(funcTable)

func (st *SymbolTable) get(ident string) symbol {
	symbol, ok := (*st)[ident]
	if !ok {
		errorf("Undefined variable '%s'", ident)
	}
	if symbol.stype == NONE {
		errorf("Uninitialized variable '%s'", ident)
	}
	return symbol
}

func (st *SymbolTable) set(ident string, symbol symbol) {
	(*st)[ident] = symbol
}

func (ft *funcTable) get(ident string) FuncDec {
	funcDec, ok := (*ft)[ident]
	if !ok {
		errorf("Undefined function '%s'", ident)
	}
	return funcDec
}

func (ft *funcTable) set(ident string, funcDec FuncDec) {
	_, ok := (*ft)[ident]
	if ok {
		errorf("Function '%s' already exists", ident)
	}
	(*ft)[ident] = funcDec
}
