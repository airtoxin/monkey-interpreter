package object

type ObjectType string

const (
	NULL_OBJ         = "NULL"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION_OBJ"
	STRING_OBJ       = "STRING_OBJ"
	ARRAY_OBJ        = "ARRAY"
	BUILTIN_OBJ      = "BUILTIN_OBJ"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
