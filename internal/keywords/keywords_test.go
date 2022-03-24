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
	})
}
