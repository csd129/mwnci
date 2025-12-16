package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)


// Geteuid ...
func Geteuid(args ...object.Object) object.Object {
	if err := typing.Check(
		"geteuid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}
	return &object.Integer{Value: int64(os.Geteuid())}
}

