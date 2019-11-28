package object

type String struct {
	Value string
}

func (i *String) Inspect() string {
	return i.Value
}

func (i *String) Type() ObjectType {
	return STRING_OBJ
}
