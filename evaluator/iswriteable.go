package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func IsWriteable(args ...object.Object) object.Object {
	if err := typing.Check(
		"iswriteable", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	file := args[0].(*object.String).Value
	f, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		return FALSE
	}
	defer f.Close()
	return TRUE
}
