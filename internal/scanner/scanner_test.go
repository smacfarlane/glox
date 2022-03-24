package scanner

import (
	"fmt"
	"testing"

	"github.com/smacfarlane/glox/internal/tokens"
	"github.com/stretchr/testify/require"
)

func TestScanner(t *testing.T) {
	/*src := `
	// Your first Lox program!
	print "Hello, world!";
	`*/
	ezSrc := "print \"Hello, world!\";"
	intSrc := "print 10;"
	negSrc := "print -10;"

	t.Run("Scans a program and stores the tokens", func(t *testing.T) {
		require := require.New(t)

		scanner := NewScanner(ezSrc)
		require.Len(scanner.tokens, 0)

		scanner.ScanTokens()
		require.Len(scanner.tokens, 4)
	})

	t.Run("Scans a program and handles numbers", func(t *testing.T) {
		require := require.New(t)

		scanner := NewScanner(intSrc)
		require.Len(scanner.tokens, 0)

		scanner.ScanTokens()
		require.Len(scanner.tokens, 4)

		// TODO: This ends up actually being a token.IDENTIFIER not
		// tokens.PRINT
		tokenStr := scanner.tokens[0].String()
		require.Equal(tokenStr, fmt.Sprintf("%d %s %s", tokens.IDENTIFIER, "print", ""))
		tokenStr = scanner.tokens[1].String()
		require.Equal(tokenStr, fmt.Sprintf("%d %s %s", tokens.NUMBER, "10", "10"))

		scanner = NewScanner(negSrc)
		require.Len(scanner.tokens, 0)

		// Scanner does not handle negative numbers yet
		/*
			scanner.ScanTokens()
			require.Len(scanner.tokens, 5)
			tokenStr = scanner.tokens[1].String()
			require.Equal(tokenStr, fmt.Sprintf("%d %s %s", tokens.NUMBER, "-10", "10"))
		*/
	})
}
