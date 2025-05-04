package parser

import (
	"github.com/RootControl/Interpreter/internal/ast"
	lx "github.com/RootControl/Interpreter/internal/lexer"
	tk "github.com/RootControl/Interpreter/internal/token"
)

type Parser struct {
	Lexer *lx.Lexer
	CurrentToken tk.Token
	PeekToken tk.Token
}

func NewParser(lexer *lx.Lexer) *Parser {
	parser := Parser{Lexer: lexer}

	parser.nextToken()
	parser.nextToken()

	return &parser
}

func (p *Parser) nextToken() {
	p.CurrentToken = p.PeekToken

	p.PeekToken, _ = p.Lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()

	for p.CurrentToken.Type != tk.EoF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.CurrentToken.Type {
	case tk.Let:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := ast.NewLetStatement(p.CurrentToken)

	if !p.expectPeek(tk.Identifier) {
		return nil
	}

	statement.Name = ast.NewIdentifier(p.CurrentToken)

	if !p.expectPeek(tk.Assign) {
		return nil
	}

	// TODO: add expression parse

	for !p.currentTokenIs(tk.Semicolon) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) currentTokenIs(tokenType tk.TokenType) bool {
	return p.CurrentToken.Type == tokenType
}

func (p *Parser) peekTokenIs(tokenType tk.TokenType) bool {
	return p.PeekToken.Type == tokenType
}

func (p *Parser) expectPeek(tokenType tk.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
