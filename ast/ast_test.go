package ast

import (
	"kisumu/token"
	"testing"
)
//tests the functionality of the expression statements
func TestExpressionStatement_statementNode(t *testing.T) {
	exprToken := token.Token{Type: token.IDENT, Literal: "foobar"}
    stmt := &ExpressionStatement{
        Token: exprToken,
    }

    if stmt.TokenLiteral() != "foobar" {
        t.Errorf("stmt.TokenLiteral() wrong. expected=%q, got=%q", "foobar", stmt.TokenLiteral())
    }
}
