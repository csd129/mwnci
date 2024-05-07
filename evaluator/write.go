package evaluator

import (
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

// Write ...
func Write(args ...object.Object) object.Object {
	if err := typing.Check(
		"write", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.INTEGER_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	fd := int(args[0].(*object.Integer).Value)
	data := []byte(args[1].(*object.String).Value)

	n, err := syscall.Write(fd, data)
	if err != nil {
		return newError("IOError: %s", err)
	}

	return &object.Integer{Value: int64(n)}
}

