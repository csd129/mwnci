package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func Printat(args ...object.Object) object.Object {
	if err := typing.Check(
		"printat", args,
		typing.MinimumArgs(3),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	X := int(args[0].(*object.Integer).Value)
	Y := int(args[1].(*object.Integer).Value)
	fmt.Print("\033[", X, ";", Y, "H")
	for _, v := range args[2:] {
		fmt.Print(v.ToInterface())
	}
	fmt.Print("\033[0m")

	return &object.Null{}
}
