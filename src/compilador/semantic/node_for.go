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
		do_s := n.Do.Eval(st)
		if do_s.stype != NONE {
			return do_s
		}
		s = n.Cond.Eval(st)
		expect(INT, s)
		n.Post.Eval(st)
	}
	return symbol{NONE, nil}
}
