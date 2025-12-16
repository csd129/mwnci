package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Trimprefix(args ...object.Object) object.Object {
	if err := typing.Check(
		"trimprefix", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	line := args[0].(*object.String).Value
	prefix := args[1].(*object.String).Value
	return &object.String{Value: strings.TrimPrefix(line, prefix)}
}

