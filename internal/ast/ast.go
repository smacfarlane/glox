package ast

import (
	"fmt"
)

type Visitor struct {
	expr Expr
}

func NewVisitor(e Expr) Visitor {
	return Visitor{expr: e}
}

func (v Visitor) String() string {
	return v.expr.Accept(v)
}

func (v Visitor) VisitBinaryExpression(expr Binary) string {
	return v.parenthesize(expr.Operator.GetLexeme(), expr.Left, expr.Right)
}

func (v Visitor) VisitGroupingExpression(expr Grouping) string {
	return v.parenthesize("group", expr.Expression)
}

func (v Visitor) VisitLiteralExpression(expr Literal) string {
	if expr.Value == "" {
		return "nil"
	}
	return expr.Value
}

func (v Visitor) VisitUnaryExpression(expr Unary) string {
	return v.parenthesize(expr.Operator.GetLexeme(), expr.Right)
}

func (v Visitor) parenthesize(name string, exprs ...Expr) string {
	value := "(" + name

	for _, expr := range exprs {
		value += fmt.Sprintf(" %s", expr.Accept(v))
	}

	value += ")"

	return value
}
