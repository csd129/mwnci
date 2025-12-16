package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"runtime"
)

func GetOs(args ...object.Object) object.Object {
	if err := typing.Check(
		"getos", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}
	return &object.String{Value: runtime.GOOS}
}

