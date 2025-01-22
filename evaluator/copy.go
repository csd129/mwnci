package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func acopyFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"copy", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	newArray := arr.Copy()
	return newArray
}
