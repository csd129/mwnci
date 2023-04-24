package evaluator

import (
	"os"
	"mwnci/object"
	"mwnci/typing"
)

// Getuid ...
func Getuid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getuid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	pid := os.Getuid()
	return &object.Integer{Value: int64(pid)}
}

func init() {
	RegisterBuiltin("getuid",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Getuid(args...))
		})
}

