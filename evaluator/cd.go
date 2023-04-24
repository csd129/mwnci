package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func Cd(args ...object.Object) object.Object {
	dir := os.Getenv("HOME")
	if err := typing.Check(
		"cd", args,
		typing.RangeOfArgs(0, 1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	if len(args) == 1 {
		dir = args[0].(*object.String).Value
	}
	err := os.Chdir(dir)
	if err != nil {
		return newError(err.Error())
	}
	return TRUE
}

func init() {
	RegisterBuiltin("cd",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Cd(args...))
		})
}
