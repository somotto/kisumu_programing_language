package ast

import "kisumu/token"

//for the base interface for all AST nodes.
type Node interface {
    TokenLiteral() string
}

// represents a statement in the language.
type Statement interface {
    Node
    statementNode()
}

// represents an expression in the language.
type Expression interface {
    Node
    expressionNode()
}

// the root node of the entire program, containing all statements.
type Program struct {
    Statements []Statement
}

// to return the literal value of the token for debugging.
func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    }
    return ""
}

// represents a variable declaration.
type LetStatement struct {
    Token token.Token // the token.LET token
    Name  *Identifier // variable name
    Value Expression  // expression assigned to the variable
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// represents a variable or function name.
type Identifier struct {
    Token token.Token // the token.IDENT token
    Value string      // the name of the identifier
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// for return statements
type ReturnStatement struct {
    Token       token.Token // the 'return' token
    ReturnValue Expression   // the returned expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// for statements containing expressions
type ExpressionStatement struct {
    Token      token.Token // first token of the expression
    Expression Expression  // expression itself
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
