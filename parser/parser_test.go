package parser

import (
    "kisumu/ast"
    "kisumu/token"
    "testing"
)

// Helper function to create tokens for testing
func createTestTokens() []token.Token {
    return []token.Token{
        {Type: token.LET, Literal: "let"},
        {Type: token.IDENT, Literal: "x"},
        {Type: token.ASSIGN, Literal: "="},
        {Type: token.INT, Literal: "5"},
        {Type: token.SEMICOLON, Literal: ";"},
    }
}

// Test the Parse function
func TestParse(t *testing.T) {
    tokens := createTestTokens()
    p := New(tokens)
    program := p.Parse()

    if len(program.Statements) != 1 {
        t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
    }

    stmt, ok := program.Statements[0].(*ast.LetStatement)
    if !ok {
        t.Fatalf("Expected *ast.LetStatement, got %T", program.Statements[0])
    }

    if stmt.Name.Value != "x" {
        t.Fatalf("Expected identifier to be 'x', got %s", stmt.Name.Value)
    }

    if stmt.Value.TokenLiteral() != "5" {
        t.Fatalf("Expected value to be '5', got %s", stmt.Value.TokenLiteral())
    }
}

// Test the parseLetStatement function
func TestParseLetStatement(t *testing.T) {
    tokens := createTestTokens()
    p := New(tokens)
    stmt := p.parseLetStatement()

    if stmt.Name.Value != "x" {
        t.Fatalf("Expected name to be 'x', got %s", stmt.Name.Value)
    }

    if stmt.Value.TokenLiteral() != "5" {
        t.Fatalf("Expected value to be '5', got %s", stmt.Value.TokenLiteral())
    }
}

// Test the parseReturnStatement function
func TestParseReturnStatement(t *testing.T) {
    tokens := []token.Token{
        {Type: token.RETURN, Literal: "return"},
        {Type: token.INT, Literal: "10"},
        {Type: token.SEMICOLON, Literal: ";"},
    }
    p := New(tokens)
    stmt := p.parseReturnStatement()

    if stmt.ReturnValue.TokenLiteral() != "10" {
        t.Fatalf("Expected return value to be '10', got %s", stmt.ReturnValue.TokenLiteral())
    }
}

// Test the parseExpression function (e.g., for integer literals)
func TestParseIntegerLiteral(t *testing.T) {
    tokens := []token.Token{
        {Type: token.INT, Literal: "42"},
    }
    p := New(tokens)
    expr := p.parseExpression()

    intLiteral, ok := expr.(*ast.IntegerLiteral)
    if !ok {
        t.Fatalf("Expected *ast.IntegerLiteral, got %T", expr)
    }

    if intLiteral.Value != "42" {
        t.Fatalf("Expected integer value to be '42', got %s", intLiteral.Value)
    }
}

// Test the parseArrayLiteral function
func TestParseArrayLiteral(t *testing.T) {
    tokens := []token.Token{
        {Type: token.LBRACKET, Literal: "["},
        {Type: token.INT, Literal: "1"},
        {Type: token.COMMA, Literal: ","},
        {Type: token.INT, Literal: "2"},
        {Type: token.RBRACKET, Literal: "]"},
    }
    p := New(tokens)
    array := p.parseArrayLiteral()

    if len(array.Elements) != 2 {
        t.Fatalf("Expected 2 elements, got %d", len(array.Elements))
    }
}
