package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Getegid ...
func Getegid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getegid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}
	return &object.Integer{Value: int64(os.Getegid())}
}
