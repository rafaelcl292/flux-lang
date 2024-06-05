package semantic

type If struct {
	Cond Node
	Then Block
	Else Block
}

func (n If) Eval(st *SymbolTable) symbol {
	s := n.Cond.Eval(st)
	expect(INT, s)
	if s.val.(int) != 0 {
		if_s := n.Then.Eval(st)
		if if_s.stype != NONE {
			return if_s
		}
	} else {
		else_s := n.Else.Eval(st)
		if else_s.stype != NONE {
			return else_s
		}
	}
	return symbol{NONE, nil}
}
