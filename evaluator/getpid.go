package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func getPidFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"getpid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	return &object.Integer{Value: int64(os.Getpid())}
}

func init() {
	RegisterBuiltin("getpid",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (getPidFun(args...))
		})
}
