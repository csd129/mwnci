package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
	"syscall"
)

func Mkfifo(args ...object.Object) object.Object {
	if err := typing.Check(
		"mkfifo", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	path := args[0].(*object.String).Value
	err := syscall.Mkfifo(path, 0600)
	if err != nil && !os.IsExist(err) {
		return FALSE
	}
	return TRUE
}

