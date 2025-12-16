package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Pop ...
func Pop(args ...object.Object) object.Object {
	if err := typing.Check(
		"pop", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	if length == 0 {
		return newError("IndexError: pop from an empty array")
	}

	element := arr.Elements[length-1]
	arr.Elements = arr.Elements[:length-1]

	return element
}

