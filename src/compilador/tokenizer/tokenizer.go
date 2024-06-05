package tokenizer

import (
	. "compiler/tokens"
	"os"
	"unicode"
	"unicode/utf8"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Tokenizer struct {
	input string
	ch    rune
	Next  Token
}

func (t *Tokenizer) scan() {
	if len(t.input) == 0 {
		t.ch = 0
		return
	}
	r, size := utf8.DecodeRuneInString(t.input)
	t.input = t.input[size:]
	t.ch = r
}

func (t *Tokenizer) readNumber() {
	var number string
	for unicode.IsDigit(t.ch) {
		number += string(t.ch)
		t.scan()
	}

	if unicode.IsLetter(t.ch) {
		println("Tokenizing error: illegal number " + number + string(t.ch))
		os.Exit(1)
		return
	}

	t.Next = Token{Type: INTEGER, Literal: number}
}

func (t *Tokenizer) readIdentifier() {
	var identifier string
	for unicode.IsLetter(t.ch) || unicode.IsDigit(t.ch) || t.ch == '_' {
		identifier += string(t.ch)
		t.scan()
	}

	switch identifier {
	case "println":
		t.Next = Token{Type: PRINT, Literal: "print"}
	case "read":
		t.Next = Token{Type: READ, Literal: "read"}
	case "if":
		t.Next = Token{Type: IF, Literal: "if"}
	case "else":
		t.Next = Token{Type: ELSE, Literal: "else"}
	case "for":
		t.Next = Token{Type: FOR, Literal: "for"}
	case "or":
		t.Next = Token{Type: OR, Literal: "or"}
	case "and":
		t.Next = Token{Type: AND, Literal: "and"}
	case "not":
		t.Next = Token{Type: NOT, Literal: "not"}
	case "return":
		t.Next = Token{Type: RETURN, Literal: "return"}
	default:
		t.Next = Token{Type: VARIABLE, Literal: identifier}
	}
}

func (t *Tokenizer) readString() {
	var str string
	t.scan()
	for t.ch != '"' {
		str += string(t.ch)
		t.scan()

		if t.ch == 0 {
			println("Tokenizing error: illegal string " + str)
			os.Exit(1)
		}
	}
	t.Next = Token{Type: STRING, Literal: str}
}

func (t *Tokenizer) NextToken() {
	for unicode.IsSpace(t.ch) {
		t.scan()
	}

	switch t.ch {
	case '+':
		t.Next = Token{Type: PLUS, Literal: "+"}
	case '-':
		t.scan()
		if t.ch == '>' {
			t.scan()
			t.Next = Token{Type: ARROW, Literal: "->"}
			return
		}
		t.Next = Token{Type: MINUS, Literal: "-"}
		return
	case '*':
		t.Next = Token{Type: MULTIPLY, Literal: "*"}
	case '/':
		t.Next = Token{Type: DIVIDE, Literal: "/"}
	case '(':
		t.Next = Token{Type: LPAREN, Literal: "("}
	case ')':
		t.Next = Token{Type: RPAREN, Literal: ")"}
	case '{':
		t.Next = Token{Type: LBRACE, Literal: "{"}
	case '}':
		t.Next = Token{Type: RBRACE, Literal: "}"}
	case '=':
		t.scan()
		if t.ch == '=' {
			t.scan()
			t.Next = Token{Type: EQUALITY, Literal: "=="}
			return
		}
		t.Next = Token{Type: EQUALS, Literal: "="}
		return
	case '<':
		t.Next = Token{Type: LESS, Literal: "<"}
	case '>':
		t.Next = Token{Type: GREATER, Literal: ">"}
	case '.':
		t.scan()
		if t.ch == '.' {
			t.scan()
			t.Next = Token{Type: CONCAT, Literal: ".."}
			return
		}
		println("Tokenizing error: illegal character .")
		os.Exit(1)
	case ',':
		t.Next = Token{Type: COMMA, Literal: ","}
	case '"':
		t.readString()
	case ';':
		t.Next = Token{Type: SEMICOLON, Literal: ";"}
	case 0:
		t.Next = Token{Type: EOF, Literal: ""}
	default:
		if unicode.IsDigit(t.ch) {
			t.readNumber()
			return
		}
		if unicode.IsLetter(t.ch) {
			t.readIdentifier()
			return
		}
		println("Tokenizing error: illegal character " + string(t.ch))
		os.Exit(1)
	}
	t.scan()

}

func CreateTokenizer(input string) *Tokenizer {
	tok := &Tokenizer{input: input}
	tok.scan()
	tok.NextToken()
	return tok
}
