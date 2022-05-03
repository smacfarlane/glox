package tokens

import (
	"fmt"
)

type TokenType int

type Token struct {
	tokentype TokenType
	lexeme    string
	literal   string
	line      int
}

const (
	// Single-character tokens
	UNKNOWN TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func NewToken(t TokenType, lexeme, literal string, line int) Token {
	return Token{
		tokentype: t,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%d %s %s", t.tokentype, t.lexeme, t.literal)
}
