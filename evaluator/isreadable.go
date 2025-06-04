package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func IsReadable(args ...object.Object) object.Object {
	if err := typing.Check(
		"isreadable", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	file := args[0].(*object.String).Value

	ReadFile, err := os.Open(file)
        defer ReadFile.Close()
	if err != nil {
		return FALSE
	}
	return TRUE
}
