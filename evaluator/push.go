package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// push something onto an array
func fpushFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"push", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
//	newArray := arr.Copy()
	arr.Append(args[1])
	return arr
}
