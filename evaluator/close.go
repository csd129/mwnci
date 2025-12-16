package evaluator

import (
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

// Close ...
func Close(args ...object.Object) object.Object {
	if err := typing.Check(
		"close", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	fd := int(args[0].(*object.Integer).Value)

	err := syscall.Close(fd)
	if err != nil {
		return newError("IOError: %s", err)
	}

	return NULL
}
