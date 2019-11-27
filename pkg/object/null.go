package object

type Null struct{}

var NULL = &Null{}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}
