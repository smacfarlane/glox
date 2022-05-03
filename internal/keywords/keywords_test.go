package keywords

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smacfarlane/glox/internal/tokens"
)

func TestKeyword(t *testing.T) {
	t.Run("Keyword returns a type if exists", func(t *testing.T) {
		require := require.New(t)

		tokenKey, ok := Keyword("and")
		require.Equal(tokenKey, tokens.AND)
		require.True(ok)

		tokenKey, ok = Keyword("class")
		require.Equal(tokenKey, tokens.CLASS)
		require.True(ok)

		tokenKey, ok = Keyword("else")
		require.Equal(tokenKey, tokens.ELSE)
		require.True(ok)
	})

	t.Run("Keyword returns false if key not exist", func(t *testing.T) {
		require := require.New(t)

		tokenKey, ok := Keyword("nope")
		require.Equal(tokenKey, tokens.UNKNOWN)
		require.False(ok)
	})
}
