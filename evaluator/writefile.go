package evaluator

import (
	"os"

	"mwnci/object"
	"mwnci/typing"
)

// WriteFile ...
func WriteFile(args ...object.Object) object.Object {
	if err := typing.Check(
		"writefile", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	data := []byte(args[1].(*object.String).Value)

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return newError("IOError: error writing file %s: %s", filename, err)
	}

	return NULL
}
