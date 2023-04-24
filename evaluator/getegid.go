package evaluator

import (
	"os"
	"mwnci/object"
	"mwnci/typing"
)

// Getegid ...
func Getegid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getegid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	pid := os.Getegid()
	return &object.Integer{Value: int64(pid)}
}

func init() {
	RegisterBuiltin("getegid",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Getegid(args...))
		})
}
