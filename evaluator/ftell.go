package evaluator

import (
       "syscall"
	"mwnci/object"
	"mwnci/typing"
)

func Ftell(args ...object.Object) object.Object {
	if err := typing.Check(
		"ftell", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	var (
		fd     int
	)

	fd = int(args[0].(*object.Integer).Value)

	offset, err := syscall.Seek(fd, 0, 1)
	if err != nil {
		return newError("IOError: %s", err)
	}

	return &object.Integer{Value: offset}
}

