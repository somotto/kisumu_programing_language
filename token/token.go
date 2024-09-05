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
	"func":        FUNCTION,
	"let":         LET,
	"if":          IF,
	"else":        ELSE,
	"return":      RETURN,
	"true":        TRUE,
	"false":       FALSE,
	"null":        NULL,
	"break":       BREAK,
	"case":        CASE,
	"chan":        CHAN,
	"const":       CONST,
	"default":     DEFAULT,
	"defer":       DEFER,
	"fallthrough": FALLTHROUGH,
	"go":          GO,
	"goto":        GOTO,
	"interface":   INTERFACE,
	"map":         MAP,
	"package":     PACKAGE,
	"range":       RANGE,
	"select":      SELECT,
	"struct":      STRUCT,
	"switch":      SWITCH,
	"type":        TYPE,
	"continue":    CONTINUE,
	"for":         FOR,
	"import":      IMPORT,
	"var":         VAR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
