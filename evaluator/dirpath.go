package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"path"
)

// Dirpath ...
func Dirpath(args ...object.Object) object.Object {
	if err := typing.Check(
		"dirpath", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	return &object.String{Value: path.Dir(args[0].(*object.String).Value)}
}
