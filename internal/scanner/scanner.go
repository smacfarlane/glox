package scanner

import (
	"fmt"
	"strconv"

	"github.com/smacfarlane/glox/internal/keywords"
	"github.com/smacfarlane/glox/internal/tokens"
)

type Scanner struct {
	source  string
	tokens  []tokens.Token
	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return Scanner{source: source, start: 0, current: 0, line: 1}
}

func (s *Scanner) ScanTokens() {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, tokens.NewToken(tokens.EOF, "", "", s.line))

}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case "(":
		s.addToken(tokens.LEFT_PAREN, "") // Does "" need special meaning? Should this be a *string
	case ")":
		s.addToken(tokens.RIGHT_PAREN, "") // Does "" need special meaning? Should this be a *string
	case "{":
		s.addToken(tokens.LEFT_BRACE, "") // Does "" need special meaning? Should this be a *string
	case "}":
		s.addToken(tokens.RIGHT_BRACE, "") // Does "" need special meaning? Should this be a *string
	case ",":
		s.addToken(tokens.COMMA, "") // Does "" need special meaning? Should this be a *string
	case ".":
		s.addToken(tokens.DOT, "") // Does "" need special meaning? Should this be a *string
	case "-":
		s.addToken(tokens.MINUS, "") // Does "" need special meaning? Should this be a *string
	case "+":
		s.addToken(tokens.PLUS, "") // Does "" need special meaning? Should this be a *string
	case ";":
		s.addToken(tokens.SEMICOLON, "") // Does "" need special meaning? Should this be a *string
	case "*":
		s.addToken(tokens.STAR, "") // Does "" need special meaning? Should this be a *string
	case "!":
		if s.peek() == "=" {
			s.advance()
			s.addToken(tokens.BANG_EQUAL, "")
		} else {
			s.addToken(tokens.BANG, "")
		}
	case "=":
		if s.peek() == "=" {
			s.advance()
			s.addToken(tokens.EQUAL_EQUAL, "")
		} else {
			s.addToken(tokens.EQUAL, "")
		}
	case "<":
		if s.peek() == "=" {
			s.advance()
			s.addToken(tokens.LESS_EQUAL, "")
		} else {
			s.addToken(tokens.LESS, "")
		}
	case ">":
		if s.peek() == "=" {
			s.advance()
			s.addToken(tokens.GREATER_EQUAL, "")
		} else {
			s.addToken(tokens.GREATER, "")
		}
	case "/":
		if s.peek() == "/" {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(tokens.SLASH, "")
		}
	case " ":
	case "\r":
	case "\t":
	case "\n":
		s.line = s.line + 1
	case `"`:
		s.handleString()
	default:
		if isDigit(c) {
			s.handleNumber()
		} else if isAlpha(c) {
			s.handleIdentifer()
		} else {
			// TODO: Handle Error
			fmt.Errorf("Unexpected character. Line %v", s.line)
		}
	}
}

func (s *Scanner) advance() string {
	c := string([]rune(s.source)[s.current])
	s.current = s.current + 1
	return c
}

func (s *Scanner) addToken(t tokens.TokenType, literal string) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, tokens.NewToken(t, text, literal, s.line))
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return ""
	}

	return string([]rune(s.source)[s.current])
}

func (s *Scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return ""
	}

	return string([]rune(s.source)[s.current+1])
}

func (s *Scanner) handleString() {
	for s.peek() != `"` && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line = s.line + 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		// TODO Handle Error
		fmt.Errorf("Unterminated String. Line %v", s.line)
	}

	s.advance()

	str := string([]rune(s.source)[s.start+1 : s.current-1])
	s.addToken(tokens.STRING, str)
}

func (s *Scanner) handleNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == "." && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	// TODO: The book adds this as a double, using an Object to hold the value.
	//       Handle conversion on access?
	number := string([]rune(s.source)[s.start:s.current])
	s.addToken(tokens.NUMBER, number)
}

func (s *Scanner) handleIdentifer() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := string([]rune(s.source)[s.start:s.current])
	if _, ok := keywords.Keyword(text); ok {
		s.addToken(tokens.IDENTIFIER, "")
	}

}

func isDigit(s string) bool {
	// Ignore error. Cannot convert means it is not a digit
	n, _ := strconv.Atoi(s)
	if n >= 0 && n <= 9 {
		return true
	}

	return false
}

func isAlpha(s string) bool {
	if (s >= "a" && s <= "z") || (s >= "A" && s <= "Z") || s == "_" {
		return true
	}
	return false
}

func isAlphaNumeric(s string) bool {
	if isAlpha(s) || isDigit(s) {
		return true
	}
	return false
}
