package ast

import (
    "testing"
    "kisumu/token"
)

func TestProgramTokenLiteral(t *testing.T) {
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
                Value: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: "5"},
            },
        },
    }

    if program.TokenLiteral() != "let" {
        t.Errorf("program.TokenLiteral() wrong. got=%q", program.TokenLiteral())
    }
}

func TestLetStatementTokenLiteral(t *testing.T) {
    letStmt := &LetStatement{
        Token: token.Token{Type: token.LET, Literal: "let"},
        Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
        Value: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: "5"},
    }

    if letStmt.TokenLiteral() != "let" {
        t.Errorf("letStmt.TokenLiteral() wrong. got=%q", letStmt.TokenLiteral())
    }
}

func TestIdentifierTokenLiteral(t *testing.T) {
    ident := &Identifier{
        Token: token.Token{Type: token.IDENT, Literal: "x"},
        Value: "x",
    }

    if ident.TokenLiteral() != "x" {
        t.Errorf("ident.TokenLiteral() wrong. got=%q", ident.TokenLiteral())
    }
}

func TestIntegerLiteralTokenLiteral(t *testing.T) {
    intLit := &IntegerLiteral{
        Token: token.Token{Type: token.INT, Literal: "5"},
        Value: "5",
    }

    if intLit.TokenLiteral() != "5" {
        t.Errorf("intLit.TokenLiteral() wrong. got=%q", intLit.TokenLiteral())
    }
}

func TestReturnStatementTokenLiteral(t *testing.T) {
    returnStmt := &ReturnStatement{
        Token: token.Token{Type: token.RETURN, Literal: "return"},
        ReturnValue: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: "5"},
    }

    if returnStmt.TokenLiteral() != "return" {
        t.Errorf("returnStmt.TokenLiteral() wrong. got=%q", returnStmt.TokenLiteral())
    }
}

func TestExpressionStatementTokenLiteral(t *testing.T) {
    exprStmt := &ExpressionStatement{
        Token:      token.Token{Type: token.INT, Literal: "5"},
        Expression: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: "5"},
    }

    if exprStmt.TokenLiteral() != "5" {
        t.Errorf("exprStmt.TokenLiteral() wrong. got=%q", exprStmt.TokenLiteral())
    }
}

func TestArrayLiteralTokenLiteral(t *testing.T) {
    arrayLit := &ArrayLiteral{
        Token:    token.Token{Type: token.LBRACKET, Literal: "["},
        Elements: []Expression{&IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: "5"}},
    }

    if arrayLit.TokenLiteral() != "[" {
        t.Errorf("arrayLit.TokenLiteral() wrong. got=%q", arrayLit.TokenLiteral())
    }
}

func TestObjectLiteralTokenLiteral(t *testing.T) {
    objectLit := &ObjectLiteral{
        Token: token.Token{Type: token.LBRACE, Literal: "{"},
        Pairs: map[string]Expression{
            "key": &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: "5"},
        },
    }

    if objectLit.TokenLiteral() != "{" {
        t.Errorf("objectLit.TokenLiteral() wrong. got=%q", objectLit.TokenLiteral())
    }
}
