package ast

import (
	"github.com/smacfarlane/glox/internal/tokens"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestASTStringRepresentation(t *testing.T) {
	testExpr := Binary{
		Left: Unary{
			Operator: tokens.NewToken(tokens.MINUS, "-", "", 1),
			Right: Literal{
				Value: "123",
			},
		},
		Operator: tokens.NewToken(tokens.STAR, "*", "", 1),
		Right: Grouping{
			Expression: Literal{
				Value: "45.67",
			},
		},
	}

	t.Run("Tests Constructed AST prints correctly", func(t *testing.T) {
		require := require.New(t)

		v := NewVisitor(testExpr)

		require.Equal(v.String(), "(* (- 123) (group 45.67))")
	})

}
