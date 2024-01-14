package evaluator

import (
	"os"
	"mwnci/object"
	"mwnci/typing"
)

func getPpidFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"getppid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	return &object.Integer{Value: int64(os.Getppid())}
}

