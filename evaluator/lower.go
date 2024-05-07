package evaluator

import (
	"strings"

	"mwnci/object"
	"mwnci/typing"
)

// Lower ...
func Lower(args ...object.Object) object.Object {
	if err := typing.Check(
		"tolower", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	str := args[0].(*object.String)
	return &object.String{Value: strings.ToLower(str.Value)}
}

