package object

import "fmt"

type Boolean struct {
	Value bool
}

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
