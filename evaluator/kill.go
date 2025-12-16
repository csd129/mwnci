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
		return newError("%s", err.Error())
	}
	err := syscall.Kill(int(args[0].(*object.Integer).Value), syscall.SIGTERM)
	if err == nil {
		return TRUE
	}
	return FALSE
}
