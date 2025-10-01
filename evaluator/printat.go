package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func Printat(args ...object.Object) object.Object {
	if err := typing.Check(
		"printat", args,
		typing.MinimumArgs(2),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	Line := args[0].(*object.Integer).Value
	Column := args[1].(*object.Integer).Value
	args = args[2:]
	fmt.Printf("\033[%d;%dH", Line, Column)
	for _, arg := range args {
		fmt.Print(arg)
	}
	return NULL
}


