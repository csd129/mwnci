package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Shift ...
func Shift(args ...object.Object) object.Object {
	if err := typing.Check(
		"shift", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	amount := args[1].(*object.Integer).Value
	newArray := arr.Copy()
	length := int64(len(newArray.Elements))
	if amount < 0 {
		amount = 0
	}
	if amount > length {
		amount = length
	}
	newArray.Elements = newArray.Elements[amount:]
	return newArray
}
