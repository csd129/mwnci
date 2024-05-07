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

	return &object.String{Value: args[0].Inspect()}
}

