package object

type ObjectType string

const (
	NULL_OBJ         = "NULL"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	BUILTIN_OBJ      = "BUILTIN"
	QUOTE_OBJ        = "QUOTE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
