package parser

import (
	"compiler/semantic"
	"compiler/tokenizer"
	. "compiler/tokens"
	"strconv"
)

func Parse(tok *tokenizer.Tokenizer) semantic.Node {
	node := block(tok)
	if tok.Next.Type != EOF {
		createError("EOF", tok.Next)
	}
	return node
}

func block(tok *tokenizer.Tokenizer) semantic.Node {
	stmts := make([]semantic.Node, 0)
	for tok.Next.Type != EOF {
		stmt := statement(tok)
		stmts = append(stmts, stmt)
		expect(tok, SEMICOLON)
	}
	return &semantic.Block{Stmts: stmts}
}

func funcCall(tok *tokenizer.Tokenizer, ident string) semantic.Node {
	tok.NextToken()
	args := make([]semantic.Node, 0)
	if tok.Next.Type != RPAREN {
		args = append(args, boolExpression(tok))
		for tok.Next.Type == COMMA {
			tok.NextToken()
			args = append(args, boolExpression(tok))
		}
	}
	expect(tok, RPAREN)
	return &semantic.FuncCall{
		Name: ident,
		Args: semantic.Block{Stmts: args},
	}
}

func statement(tok *tokenizer.Tokenizer) semantic.Node {
	switch tok.Next.Type {
	case PRINT:
		tok.NextToken()
		expr := boolExpression(tok)
		return &semantic.Print{Expr: expr}
	case VARIABLE:
		ident := tok.Next.Literal
		tok.NextToken()
		if tok.Next.Type == LPAREN {
			call := funcCall(tok, ident)
			return call
		}
		if tok.Next.Type == ARROW {
			tok.NextToken()
			vars := make([]semantic.Assign, 0)
			if tok.Next.Type == VARIABLE {
				vars = append(vars, semantic.Assign{Ident: tok.Next.Literal})
				tok.NextToken()
				for tok.Next.Type == COMMA {
					tok.NextToken()
					vars = append(vars, semantic.Assign{Ident: tok.Next.Literal})
					expect(tok, VARIABLE)
				}
			}
			expect(tok, LBRACE)
			stmts := make([]semantic.Node, 0)
			for tok.Next.Type != RBRACE {
				stmt := statement(tok)
				stmts = append(stmts, stmt)
			}
			tok.NextToken()
			return &semantic.FuncDec{
				Ident:     ident,
				Vars:      vars,
				FuncBlock: semantic.Block{Stmts: stmts},
			}
		}
		expect(tok, EQUALS)
		if tok.Next.Type == LBRACE {
			tok.NextToken()
			stmts := make([]semantic.Node, 0)
			for tok.Next.Type != RBRACE {
				stmt := statement(tok)
				stmts = append(stmts, stmt)
			}
			tok.NextToken()
			return &semantic.Assign{Ident: ident, Expr: semantic.Block{Stmts: stmts}}
		}
		expr := boolExpression(tok)
		return &semantic.Assign{Ident: ident, Expr: expr}
	case IF:
		tok.NextToken()
		node := boolExpression(tok)
		expect(tok, LBRACE)
		then_stmts := make([]semantic.Node, 0)
		for tok.Next.Type != RBRACE {
			stmt := statement(tok)
			then_stmts = append(then_stmts, stmt)
		}
		else_stmts := make([]semantic.Node, 0)
		tok.NextToken()
		if tok.Next.Type == ELSE {
			tok.NextToken()
			expect(tok, LBRACE)
			for tok.Next.Type != RBRACE {
				stmt := statement(tok)
				else_stmts = append(else_stmts, stmt)
			}
			tok.NextToken()
		}
		return &semantic.If{
			Cond: node,
			Then: semantic.Block{Stmts: then_stmts},
			Else: semantic.Block{Stmts: else_stmts},
		}
	case FOR:
		tok.NextToken()
		var assign semantic.Node
		if tok.Next.Type == SEMICOLON {
			assign = &semantic.NoOp{}
		} else {
			assign = statement(tok)
		}
		expect(tok, SEMICOLON)
		var cond semantic.Node
		if tok.Next.Type == SEMICOLON {
			cond = &semantic.IntVal{Val: 1}
		} else {
			cond = boolExpression(tok)
		}
		expect(tok, SEMICOLON)
		var assign2 semantic.Node
		if tok.Next.Type == LBRACE {
			assign2 = &semantic.NoOp{}
		} else {
			assign2 = statement(tok)
		}
		expect(tok, LBRACE)
		stmts := make([]semantic.Node, 0)
		for tok.Next.Type != RBRACE {
			stmt := statement(tok)
			stmts = append(stmts, stmt)
		}
		tok.NextToken()
		return &semantic.For{
			Init: assign,
			Cond: cond,
			Post: assign2,
			Do:   semantic.Block{Stmts: stmts},
		}
	case RETURN:
		tok.NextToken()
		if tok.Next.Type == SEMICOLON {
			return &semantic.Return{Expr: &semantic.IntVal{Val: 0}}
		}
		expr := boolExpression(tok)
		return &semantic.Return{Expr: expr}
	case SEMICOLON:
		tok.NextToken()
		return &semantic.NoOp{}
	default:
		createError("STATEMENT", tok.Next)
		return nil
	}
}

func expression(tok *tokenizer.Tokenizer) semantic.Node {
	left := term(tok)
	for {
		switch tok.Next.Type {
		case PLUS, MINUS, CONCAT:
			op := tok.Next.Literal
			tok.NextToken()
			right := term(tok)
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		default:
			return left
		}
	}
}

func term(tok *tokenizer.Tokenizer) semantic.Node {
	left := factor(tok)
	for {
		if tok.Next.Type == MULTIPLY || tok.Next.Type == DIVIDE {
			op := tok.Next.Literal
			tok.NextToken()
			right := factor(tok)
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		} else {
			return left
		}
	}
}

func factor(tok *tokenizer.Tokenizer) semantic.Node {
	switch tok.Next.Type {
	case INTEGER:
		value, _ := strconv.Atoi(tok.Next.Literal)
		tok.NextToken()
		return &semantic.IntVal{Val: value}
	case STRING:
		value := tok.Next.Literal
		tok.NextToken()
		return &semantic.StrVal{Val: value}
	case PLUS, MINUS, NOT:
		op := tok.Next.Literal
		tok.NextToken()
		node := factor(tok)
		return &semantic.UnOp{Op: op, Expr: node}
	case VARIABLE:
		name := tok.Next.Literal
		tok.NextToken()
		if tok.Next.Type == LPAREN {
			return funcCall(tok, name)
		}
		return &semantic.Ident{Name: name}
	case LPAREN:
		tok.NextToken()
		node := boolExpression(tok)
		expect(tok, RPAREN)
		return node
	case READ:
		tok.NextToken()
		expect(tok, LPAREN)
		expect(tok, RPAREN)
		return &semantic.Read{}
	default:
		createError("EXPRESSION", tok.Next)
		return nil
	}
}

func boolExpression(tok *tokenizer.Tokenizer) semantic.Node {
	left := boolTerm(tok)
	for {
		if tok.Next.Type == OR {
			op := tok.Next.Literal
			tok.NextToken()
			right := boolTerm(tok)
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		} else {
			return left
		}
	}
}

func boolTerm(tok *tokenizer.Tokenizer) semantic.Node {
	left := relExpr(tok)
	for {
		if tok.Next.Type == AND {
			op := tok.Next.Literal
			tok.NextToken()
			right := relExpr(tok)
			left = &semantic.BinOp{Op: op, Left: left, Right: right}
		} else {
			return left
		}
	}
}

func relExpr(tok *tokenizer.Tokenizer) semantic.Node {
	left := expression(tok)
	switch tok.Next.Type {
	case LESS, GREATER, EQUALITY:
		op := tok.Next.Literal
		tok.NextToken()
		right := expression(tok)
		return &semantic.BinOp{Op: op, Left: left, Right: right}
	}
	return left
}
