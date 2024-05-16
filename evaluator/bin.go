package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"strconv"
)

// Bin ...
func Bin(args ...object.Object) object.Object {
	if err := typing.Check(
		"bin", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	i := args[0].(*object.Integer)
	return &object.String{Value: fmt.Sprintf("0b%s", strconv.FormatInt(i.Value, 2))}
}
