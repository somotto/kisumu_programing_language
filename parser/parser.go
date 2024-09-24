package parser

import (
    "fmt"
    "kisumu/ast"
    "kisumu/token"
)

// state of the parser
type Parser struct {
    tokens  []token.Token
    current int
    errors   []string
}

// creating a new parser instance
func New(tokens []token.Token) *Parser {
    return &Parser{tokens: tokens, current: 0}
}

// parsing the token stream and returning an AST Program
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

// parsing a single statement
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

// parsing a let statement
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

// parsing a return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    p.consume(token.RETURN) // Consume 'return' token
    value := p.parseExpression() // Parse the return value

    return &ast.ReturnStatement{
        Token:       token.Token{Type: token.RETURN, Literal: "return"},
        ReturnValue: value,
    }
}

// parsing an expression statement
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
    expression := p.parseExpression() // Parse the expression
    return &ast.ExpressionStatement{
        Token:      token.Token{Type: token.IDENT, Literal: expression.TokenLiteral()},
        Expression: expression,
    }
}

// handles parsing for various expression types
func (p *Parser) parseExpression() ast.Expression {
    switch p.peek().Type {
    case token.IDENT:
        return p.parseIdentifier()
    case token.INT:
        return p.parseIntegerLiteral()
    case token.LBRACKET:
        return p.parseArrayLiteral()
    case token.LBRACE:
        return p.parseObjectLiteral()
    default:
        return nil // handle unexpected types
    }
}

// parsing an identifier
func (p *Parser) parseIdentifier() *ast.Identifier {
    ident := p.consume(token.IDENT)
    return &ast.Identifier{
        Token: token.Token{Type: token.IDENT, Literal: ident.Literal},
        Value: ident.Literal,
    }
}

// parsing an integer literal
func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
    intToken := p.consume(token.INT)
    return &ast.IntegerLiteral{
        Token: token.Token{Type: token.INT, Literal: intToken.Literal},
        Value: intToken.Literal, // convert to appropriate type if necessary
    }
}

// parsing an array literal
func (p *Parser) parseArrayLiteral() *ast.ArrayLiteral {
    p.consume(token.LBRACKET) // Consume '['
    elements := []ast.Expression{}

    for !p.isAtEnd() && p.peek().Type != token.RBRACKET {
        element := p.parseExpression() // Parse each element
        elements = append(elements, element)

        if p.peek().Type == token.COMMA {
            p.consume(token.COMMA) // to consume the comma
        } else {
            break
        }
    }

    p.consume(token.RBRACKET) // for ']'
    return &ast.ArrayLiteral{
        Token:    token.Token{Type: token.LBRACKET, Literal: "["},
        Elements: elements,
    }
}

// parsing an object literal
func (p *Parser) parseObjectLiteral() *ast.ObjectLiteral {
    p.consume(token.LBRACE) // for '{'
    pairs := make(map[string]ast.Expression)

    for !p.isAtEnd() && p.peek().Type != token.RBRACE {
        key := p.consume(token.IDENT).Literal
        p.consume(token.COLON) // for ':'
        value := p.parseExpression() // Parse the value

        pairs[key] = value

        if p.peek().Type == token.COMMA {
            p.consume(token.COMMA) // for comma
        } else {
            break
        }
    }

    p.consume(token.RBRACE) // for '}'
    return &ast.ObjectLiteral{
        Token: token.Token{Type: token.LBRACE, Literal: "{"},
        Pairs: pairs,
    }
}

// advances the current token and returns it, checking if it matches the expected type
func (p *Parser) consume(expected token.TokenType) token.Token {
    if p.peek().Type != expected {
        panic(fmt.Sprintf("Expected token %s, got %s", expected, p.peek().Type))
    }
    token := p.tokens[p.current]
    p.current++
    return token
}

// returns the current token without advancing
func (p *Parser) peek() token.Token {
    return p.tokens[p.current]
}

// checking if the parser has reached the end of the token stream
func (p *Parser) isAtEnd() bool {
    return p.current >= len(p.tokens)
}
