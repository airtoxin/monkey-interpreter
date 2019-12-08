package object

import "monkey-interpreter/pkg/ast"

type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() ObjectType {
	return QUOTE_OBJ
}

func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}
