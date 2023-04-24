package evaluator

import (
	"os"
	"mwnci/object"
	"mwnci/typing"
)

// Getgid ...
func Getgid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getgid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}

	pid := os.Getgid()
	return &object.Integer{Value: int64(pid)}
}

func init() {
	RegisterBuiltin("getgid",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Getgid(args...))
		})
}
