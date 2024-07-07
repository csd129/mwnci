package evaluator

import (
	"fmt"
	"strconv"

	"mwnci/object"
	"mwnci/typing"
)

func Bin2dec(args ...object.Object) object.Object {
	if err := typing.Check(
		"bin2dec", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	Number:=fmt.Sprintf("%d", (args[0].(*object.Integer).Value))
	i, err := strconv.ParseInt(Number, 2, 64)
	if err != nil {
		return newError(err.Error())
	}
	return &object.Integer{Value: i}
}
