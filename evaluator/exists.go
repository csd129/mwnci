package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// File ...
func Exists(args ...object.Object) object.Object {
	if err := typing.Check(
		"exists", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	file := args[0].(*object.String).Value

	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return FALSE
	} else {
		return TRUE
	}
}

func init() {
	RegisterBuiltin("exists",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Exists(args...))
		})
}
