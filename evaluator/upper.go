package evaluator

import (
	"strings"

	"mwnci/object"
	"mwnci/typing"
)

// Upper ...
func Upper(args ...object.Object) object.Object {
	if err := typing.Check(
		"toupper", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	str := args[0].(*object.String)
	return &object.String{Value: strings.ToUpper(str.Value)}
}

