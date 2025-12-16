package evaluator

import (
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

// Seek ...
func Seek(args ...object.Object) object.Object {
	if err := typing.Check(
		"seek", args,
		typing.RangeOfArgs(1, 3),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	var (
		fd     int
		whence = 0
	)

	fd = int(args[0].(*object.Integer).Value)
	offset := args[1].(*object.Integer).Value

	if len(args) == 3 {
		whence = int(args[2].(*object.Integer).Value)
	}

	offset, err := syscall.Seek(fd, offset, whence)
	if err != nil {
		return newError("IOError: %s", err)
	}

	return &object.Integer{Value: offset}
}

