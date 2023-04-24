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
	amount :=  args[1].(*object.Integer).Value
	newArray := arr.Copy()
	i:=int64(1)
	for (i <= amount) {
		newArray.PopLeft()
		i++
	}
	return newArray
}
func init() {
	RegisterBuiltin("shift",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Shift(args...))
		})
}
