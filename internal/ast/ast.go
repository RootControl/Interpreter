package ast

import (
	tk "github.com/RootControl/Interpreter/internal/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}



type Program struct {
	Statements []Statement
}

func NewProgram() *Program {
	return &Program{
		Statements: []Statement{},
	}
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}

	return p.Statements[0].TokenLiteral()
}



type LetStatement struct {
	Token tk.Token
	Name *Identifier
	Value Expression
}

func NewLetStatement(token tk.Token) *LetStatement {
	return &LetStatement{
		Token: token,
	}
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}




type Identifier struct {
	Token tk.Token							
	Value string
}

func NewIdentifier(token tk.Token) *Identifier {
	return &Identifier{
		Token: token,
		Value: token.Literal,
	}
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
