package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Len ...
func FunLen(args ...object.Object) object.Object {
	if err := typing.Check(
		"len", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	if size, ok := args[0].(object.Sizeable); ok {
		return &object.Integer{Value: int64(size.Len())}
	}
	return newError("TypeError: object of type '%s' has no len()", args[0].Type())
}

func init() {
	RegisterBuiltin("len",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (FunLen(args...))
		})
}
