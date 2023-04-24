package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Rest ...
func Rest(args ...object.Object) object.Object {
	if err := typing.Check(
		"rest", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	newArray := arr.Copy()
	newArray.PopLeft()
	return newArray
}
func init() {
	RegisterBuiltin("rest",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Rest(args...))
		})
}
