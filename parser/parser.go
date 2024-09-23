package parser

import (
    "kisumu/ast"
    "kisumu/token"
)

// state of the parser
type Parser struct {
    tokens  []token.Token
    current int
}

// creating new parser instance
func New(tokens []token.Token) *Parser {
    return &Parser{tokens: tokens, current: 0}
}

//parsing the token stream and returns an AST Program
func (p *Parser) Parse() *ast.Program {
    program := &ast.Program{}
    for !p.isAtEnd() {
        statement := p.parseStatement()
        if statement != nil {
            program.Statements = append(program.Statements, statement)
        }
    }
    return program
}

//  parseing a single statement
func (p *Parser) parseStatement() ast.Statement {
    switch p.peek().Type {
    case token.LET:
        return p.parseLetStatement()
    case token.RETURN:
        return p.parseReturnStatement()
    default:
        return p.parseExpressionStatement()
    }
}

// for parsing a let statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
    p.consume(token.LET) // Consume 'let' token
    identifier := p.consume(token.IDENT) // Consume identifier

    p.consume(token.ASSIGN) // Consume '=' token
    value := p.parseExpression() // Parse the expression

    return &ast.LetStatement{
        Token: token.Token{Type: token.LET, Literal: "let"},
        Name:  &ast.Identifier{Token: identifier, Value: identifier.Literal},
        Value: value,
    }
}

// for parsing a return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    p.consume(token.RETURN) // Consume 'return' token
    value := p.parseExpression() // Parse the return value

    return &ast.ReturnStatement{
        Token:       token.Token{Type: token.RETURN, Literal: "return"},
        ReturnValue: value,
    }
}

// it parses an expression statement
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
    expression := p.parseExpression() // Parse the expression
    return &ast.ExpressionStatement{
        Token:      token.Token{Type: token.IDENT, Literal: expression.TokenLiteral()},
        Expression: expression,
    }
}

//  handles parsing for various expression types
func (p *Parser) parseExpression() ast.Expression {
    if p.peek().Type == token.IDENT {
        ident := p.consume(token.IDENT) // Consume identifier token
        return &ast.Identifier{
            Token: token.Token{Type: token.IDENT, Literal: ident.Literal},
            Value: ident.Literal,
        }
    } else if p.peek().Type == token.INT {
        intToken := p.consume(token.INT)
        return &ast.Identifier{ // can be replaced with actual integer type if necessary
            Token: token.Token{Type: token.INT, Literal: intToken.Literal},
            Value: intToken.Literal,
        }
    }

    return nil // this is for an appropriate expression type
}

// advances the current token and returns it, checking if it matches the expected type
func (p *Parser) consume(expected token.TokenType) token.Token {
    if p.peek().Type != expected {
        panic("Unexpected token")
    }
    token := p.tokens[p.current]
    p.current++
    return token
}

// it returns the current token without advancing
func (p *Parser) peek() token.Token {
    return p.tokens[p.current]
}

// checking if the parser has reached the end of the token stream
func (p *Parser) isAtEnd() bool {
    return p.current >= len(p.tokens)
}
