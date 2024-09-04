package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    IDENT  = "IDENT"
    INT    = "INT"
    FLOAT  = "FLOAT"
    STRING = "STRING"
    BOOL   = "BOOL"

    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"

    LT = "<"
    GT = ">"

    EQ     = "=="
    NOT_EQ = "!="

    COMMA     = ","
    SEMICOLON = ";"
    COLON     = ":"
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"
    LBRACKET  = "["
    RBRACKET  = "]"

    FUNCTION = "FUNCTION"
    LET      = "LET"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    NULL     = "NULL"
)

var keywords = map[string]TokenType{
    "func":   FUNCTION,
    "let":    LET,
    "if":     IF,
    "else":   ELSE,
    "return": RETURN,
    "true":   TRUE,
    "false":  FALSE,
    "null":   NULL,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
