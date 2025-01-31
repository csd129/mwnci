package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// First ...
func First(args ...object.Object) object.Object {
	if err := typing.Check(
		"first", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}
	return NULL
}

