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
		return newError(err.Error())
	}
	return &object.Integer{Value: int64(os.Getuid())}
}

func init() {
	RegisterBuiltin("getuid",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Getuid(args...))
		})
}
