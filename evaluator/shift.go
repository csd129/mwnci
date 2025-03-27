package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Shift ...
func Shift(args ...object.Object) object.Object {
	if err := typing.Check(
		"shift", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	amount := 1
	arr := args[0].(*object.Array)
	if len(args) == 2 {
		amount = int(args[1].(*object.Integer).Value)
	}
	newArray := arr.Copy()
	length := len(newArray.Elements)
	if amount < 0 {
		amount = 0
	}
	if amount > length {
		amount = length
	}
	newArray.Elements = newArray.Elements[amount:]
	return newArray
}
