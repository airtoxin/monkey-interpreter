package ast

import "monkey-interpreter/pkg/token"

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (ident *Identifier) expressionNode() {}

func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Literal
}
