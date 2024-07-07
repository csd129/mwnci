package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func Getgid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getgid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}

	return &object.Integer{Value: int64(os.Getgid())}
}

