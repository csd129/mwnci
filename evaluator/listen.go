package evaluator

import (
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

// Listen ...
func Listen(args ...object.Object) object.Object {
	if err := typing.Check(
		"listen", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	fd := int(args[0].(*object.Integer).Value)
	backlog := int(args[1].(*object.Integer).Value)

	if err := syscall.Listen(fd, backlog); err != nil {
		return newError("SocketError: %s", err)
	}

	return NULL
}
