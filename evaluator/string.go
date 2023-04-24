package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func strFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"string", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	out := args[0].Inspect()
	return &object.String{Value: out}
}

func init() {
	RegisterBuiltin("string",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (strFun(args...))
		})
}
