package ast

import (
	"kisumu/token"
	"testing"
)

func TestExpressionStatement_statementNode(t *testing.T) {
	type fields struct {
		Token      token.Token
		Expression Expression
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := &ExpressionStatement{
				Token:      tt.fields.Token,
				Expression: tt.fields.Expression,
			}
			es.statementNode()
		})
	}
}
