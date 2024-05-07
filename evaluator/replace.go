package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Replace(args ...object.Object) object.Object {
	if err := typing.Check(
		"replace", args,
		typing.ExactArgs(4),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	line := args[0].(*object.String).Value
	search := args[1].(*object.String).Value
	rep := args[2].(*object.String).Value
	count := int(args[3].(*object.Integer).Value)
	return &object.String{Value: strings.Replace(line, search, rep, count)}
}

