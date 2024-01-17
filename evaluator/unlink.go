package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Remove a file/directory.
func unlinkFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"unlink", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	path := args[0].Inspect()

	err := os.Remove(path)
	if err != nil {
		return &object.Boolean{Value: false}
	}
	return &object.Boolean{Value: true}
}
