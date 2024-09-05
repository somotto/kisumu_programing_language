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
        {"break", BREAK},
        {"case", CASE},
        {"chan", CHAN},
        {"const", CONST},
        {"default", DEFAULT},
        {"defer", DEFER},
        {"fallthrough", FALLTHROUGH},
        {"go", GO},
        {"goto", GOTO},
        {"interface", INTERFACE},
        {"map", MAP},
        {"package", PACKAGE},
        {"range", RANGE},
        {"select", SELECT},
        {"struct", STRUCT},
        {"switch", SWITCH},
        {"type", TYPE},
        {"continue", CONTINUE},
        {"for", FOR},
        {"import", IMPORT},
        {"var", VAR},
        {"foobar", IDENT},   // Not a keyword, to return IDENT
        {"x", IDENT},        // Not a keyword, to return IDENT
        {"myVar", IDENT},    // Not a keyword, to return IDENT
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
        {BREAK, "break"},
        {CASE, "case"},
        {CHAN, "chan"},
        {CONST, "const"},
        {DEFAULT, "default"},
        {DEFER, "defer"},
        {FALLTHROUGH, "fallthrough"},
        {GO, "go"},
        {GOTO, "goto"},
        {INTERFACE, "interface"},
        {MAP, "map"},
        {PACKAGE, "package"},
        {RANGE, "range"},
        {SELECT, "select"},
        {STRUCT, "struct"},
        {SWITCH, "switch"},
        {TYPE, "type"},
        {CONTINUE, "continue"},
        {FOR, "for"},
        {IMPORT, "import"},
        {VAR, "var"},
    }

    for _, tt := range tests {
        tok := Token{Type: tt.tokenType, Literal: tt.literal}
        if tok.Type != tt.tokenType || tok.Literal != tt.literal {
            t.Errorf("Token creation failed for Type %q, Literal %q", tt.tokenType, tt.literal)
        }
    }
}
