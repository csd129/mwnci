package evaluator

import (
	"syscall"
	"mwnci/object"
	"mwnci/typing"
)

// Close ...
func FClose(args ...object.Object) object.Object {
	if err := typing.Check(
		"fclose", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}     

	fd := int(args[0].(*object.Integer).Value)

	err := syscall.Close(fd)
	if err != nil {
		return newError(err.Error())
	}

	return TRUE
}

func init() {
	RegisterBuiltin("fclose",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (FClose(args...))
		})
}
