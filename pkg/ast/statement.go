package ast

import "monkey-interpreter/pkg/token"

type Statement interface {
	Node
	statementNode() // dummy
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode() {}

func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}
