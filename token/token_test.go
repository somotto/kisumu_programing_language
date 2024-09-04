package token

import (
	"testing"
)

// TestLookupIdent checks that keywords are correctly identified and that identifiers are handled properly.
func TestLookupIdent(t *testing.T) {
    tests := []struct {
        input    string
        expected TokenType
    }{
        {"func", FUNCTION},
        {"let", LET},
        {"if", IF},
        {"else", ELSE},
        {"return", RETURN},
        {"true", TRUE},
        {"false", FALSE},
        {"null", NULL},
        {"foobar", IDENT},   // non-keyword identifier
        {"x", IDENT},        // non-keyword identifier
        {"myVar", IDENT},    // non-keyword identifier
    }

    for _, tt := range tests {
        tokType := LookupIdent(tt.input)
        if tokType != tt.expected {
            t.Errorf("LookupIdent(%q) = %q, want %q", tt.input, tokType, tt.expected)
        }
    }
}

// TestTokenCreation checks that tokens are created correctly.
func TestTokenCreation(t *testing.T) {
    tests := []struct {
        tokenType TokenType
        literal   string
    }{
        {ASSIGN, "="},
        {PLUS, "+"},
        {MINUS, "-"},
        {BANG, "!"},
        {ASTERISK, "*"},
        {SLASH, "/"},
        {INT, "123"},
        {FLOAT, "123.456"},
        {STRING, "hello"},
        {BOOL, "true"},
    }

    for _, tt := range tests {
        tok := Token{Type: tt.tokenType, Literal: tt.literal}
        if tok.Type != tt.tokenType || tok.Literal != tt.literal {
            t.Errorf("Token creation failed for Type %q, Literal %q", tt.tokenType, tt.literal)
        }
    }
}
