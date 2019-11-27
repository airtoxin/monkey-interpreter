package evaluator

import (
	"fmt"
	"monkey-interpreter/pkg/object"
)

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
