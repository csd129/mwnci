package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func Chown(args ...object.Object) object.Object {
	if err := typing.Check(
		"chown", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.STRING_OBJ, object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].String()
	owner := int(args[1].(*object.Integer).Value)
	group := int(args[2].(*object.Integer).Value)

	err := os.Chown(path, owner, group)
	if err != nil {
		return newError(err.Error())
	}
	return TRUE
}

