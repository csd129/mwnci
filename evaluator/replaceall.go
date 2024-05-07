package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Replaceall(args ...object.Object) object.Object {
	if err := typing.Check(
		"replaceall", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	line := args[0].(*object.String).Value
	search := args[1].(*object.String).Value
	rep := args[2].(*object.String).Value
	return &object.String{Value: strings.ReplaceAll(line, search, rep)}
}

