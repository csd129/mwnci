package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func TrimSuff(args ...object.Object) object.Object {
	if err := typing.Check(
		"trimsuffix", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	line := args[0].(*object.String).Value
	suffix := args[1].(*object.String).Value
	return &object.String{Value: strings.TrimSuffix(line, suffix)}
}

