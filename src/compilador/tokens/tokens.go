package tokens

type TokenType string

const (
	// Special tokens
	EOF       TokenType = "EOF"
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"

	// Literals
	INTEGER TokenType = "INTEGER"
	STRING  TokenType = "STRING"

	// Keywords
	PRINT    TokenType = "PRINT"
	READ     TokenType = "READ"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	FOR      TokenType = "FOR"
	OR       TokenType = "OR"
	AND      TokenType = "AND"
	NOT      TokenType = "NOT"
	RETURN   TokenType = "RETURN"

	// Operators
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE   TokenType = "DIVIDE"
	LPAREN   TokenType = "LPAREN"
	RPAREN   TokenType = "RPAREN"
	LBRACE   TokenType = "LBRACE"
	RBRACE   TokenType = "RBRACE"
	EQUALS   TokenType = "EQUALS"
	LESS     TokenType = "LESS"
	GREATER  TokenType = "GREATER"
	EQUALITY TokenType = "EQUALITY"
	CONCAT   TokenType = "CONCAT"
	ARROW    TokenType = "ARROW"

	// Variables
	VARIABLE TokenType = "VARIABLE"
)
