package evaluator

import (
	//	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Cat ...
func ConCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	data, err := os.ReadFile(args[0].(*object.String).Value)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	return &object.String{Value: string(data)}
}
