package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Getwd ...
func Getwd(args ...object.Object) object.Object {
	if err := typing.Check(
		"getwd", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	wd, err := os.Getwd()
	if err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: wd}
}
