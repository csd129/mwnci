package evaluator

import (
	"fmt"
	"strconv"

	"mwnci/object"
	"mwnci/typing"
)

func Bin2hex(args ...object.Object) object.Object {
	if err := typing.Check(
		"bin2hex", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	Number:=fmt.Sprintf("%d", (args[0].(*object.Integer).Value))
	i, err := strconv.ParseUint(Number, 2, 64)
	if err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: fmt.Sprintf("0x%x", i)}
}
