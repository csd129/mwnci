package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Getuid ...
func Getuid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getuid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}
	return &object.Integer{Value: int64(os.Getuid())}
}
