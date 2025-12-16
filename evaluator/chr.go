package evaluator

import (
	"fmt"

	"mwnci/object"
	"mwnci/typing"
)

// Chr ...
func Chr(args ...object.Object) object.Object {
	if err := typing.Check(
		"chr", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	return &object.String{Value: fmt.Sprintf("%c", rune(args[0].(*object.Integer).Value))}
}

