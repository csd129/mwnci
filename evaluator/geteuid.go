package evaluator

import (
	"os"
	"mwnci/object"
	"mwnci/typing"
)

// Geteuid ...
func Geteuid(args ...object.Object) object.Object {
	if err := typing.Check(
		"geteuid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	pid := os.Geteuid()
	return &object.Integer{Value: int64(pid)}
}

func init() {
	RegisterBuiltin("geteuid",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Geteuid(args...))
		})
}
