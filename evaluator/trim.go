package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Trim(args ...object.Object) object.Object {
	if err := typing.Check(
		"trim", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	trimmer := " \n\t\r"
	line := args[0].(*object.String).Value
	if len(args) == 2 {
		trimmer = args[1].(*object.String).Value
	}
	return &object.String{Value: strings.Trim(line, trimmer)}
}
