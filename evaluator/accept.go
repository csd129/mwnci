package evaluator

import (
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

// Accept ...
func Accept(args ...object.Object) object.Object {
	if err := typing.Check(
		"accept", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	var (
		nfd int
		err error
	)

	fd := int(args[0].(*object.Integer).Value)

	nfd, _, err = syscall.Accept(fd)
	if err != nil {
		return newError("SocketError: %s", err)
	}

	return &object.Integer{Value: int64(nfd)}
}
