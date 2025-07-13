package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Swapper(args ...object.Object) object.Object {
	if err := typing.Check(
		"swap", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	Array := args[0].(*object.Array)
//	Array := Array.Copy()
	a := args[1].(*object.Integer).Value
	b := args[2].(*object.Integer).Value
	if a < 0 || a > int64(len(Array.Elements)-1) {
		return newError("First array index (%d) out of range", a)
	}
	if b < 0 || b > int64(len(Array.Elements)-1) {
		return newError("Second array index (%d) out of range", b)
	}
	Array.Swap(int(a), int(b))
	return Array
}
