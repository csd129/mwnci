package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// push something onto an array
func acopyFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"acopy", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	newArray := arr.Copy()
	return newArray
}

func init() {
	RegisterBuiltin("acopy",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (acopyFun(args...))
		})
}
