package semantic

type For struct {
	Init Node
	Cond Node
	Post Node
	Do   Block
}

func (n For) Eval(st *SymbolTable) symbol {
	n.Init.Eval(st)
	s := n.Cond.Eval(st)
	expect(INT, s)
	for s.val.(int) != 0 {
		if n.Do.Eval(st).stype != NONE {
			break
		}
		s = n.Cond.Eval(st)
		expect(INT, s)
		n.Post.Eval(st)
	}
	return symbol{NONE, nil}
}
