package evaluator

import (
	"syscall"
	"mwnci/object"
	"mwnci/typing"
)

func Kill(args ...object.Object) object.Object {
	if err := typing.Check(
		"kill", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	p := args[0].(*object.Integer)
	P := int(p.Value)
	err := syscall.Kill(P, syscall.SIGTERM)
	if err == nil {
		return TRUE
	}
	return FALSE
}

func init() {
	RegisterBuiltin("kill",
		func(env *object.Environment, args ...object.Object) object.Object {
			return(Kill(args...))
		})
}
