package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"syscall"
)

func Kill(args ...object.Object) object.Object {
	if err := typing.Check(
		"kill", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	//P := int(args[0].(*object.Integer).Value)
	err := syscall.Kill(int(args[0].(*object.Integer).Value), syscall.SIGTERM)
	if err == nil {
		return TRUE
	}
	return FALSE
}

func init() {
	RegisterBuiltin("kill",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Kill(args...))
		})
}
