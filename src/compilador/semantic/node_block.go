package semantic

import "reflect"

type Block struct {
	Stmts []Node
}

func (n Block) Eval(st *SymbolTable) symbol {
	for _, stmt := range n.Stmts {
		typeOfStmt := reflect.TypeOf(stmt).String()
		if typeOfStmt == "*semantic.Return" {
			return stmt.Eval(st)
		}
		if  typeOfStmt == "*semantic.If" || typeOfStmt == "*semantic.For" {
			s := stmt.Eval(st)
			if s.stype == NONE {
				continue
			}
			return s
		}
		stmt.Eval(st)
	}
	return symbol{NONE, nil}
}
