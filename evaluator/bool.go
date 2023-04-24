package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Bool ...
func Bool(args ...object.Object) object.Object {
	if err := typing.Check(
		"bool", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.BOOLEAN_OBJ),
	); err != nil {
		return newError(err.Error())
	}     

	return &object.Boolean{Value: args[0].(*object.Boolean).Value}
}

func init() {
	RegisterBuiltin("bool",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Bool(args...))
		})
}
