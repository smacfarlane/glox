package keywords

import (
	"github.com/smacfarlane/glox/internal/tokens"
)

var keywords = map[string]tokens.TokenType{
	"and":    tokens.AND,
	"class":  tokens.CLASS,
	"else":   tokens.ELSE,
	"false":  tokens.FALSE,
	"for":    tokens.FOR,
	"fun":    tokens.FUN,
	"if":     tokens.IF,
	"nil":    tokens.NIL,
	"or":     tokens.OR,
	"print":  tokens.PRINT,
	"return": tokens.RETURN,
	"super":  tokens.SUPER,
	"this":   tokens.THIS,
	"true":   tokens.TRUE,
	"var":    tokens.VAR,
	"while":  tokens.WHILE,
}

func Keyword(key string) (tokens.TokenType, bool) {
	token, ok := keywords[key]
	return token, ok
}
