package evaluator

import (
	"fmt"
	"strconv"

	"mwnci/object"
	"mwnci/typing"
)

// Oct ...
func Oct(args ...object.Object) object.Object {
	if err := typing.Check(
		"oct", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	i := args[0].(*object.Integer)
	return &object.String{Value: fmt.Sprintf("0%s", strconv.FormatInt(i.Value, 8))}
}






