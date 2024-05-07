package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"strconv"
)

func IsNumber(args ...object.Object) object.Object {
	if err := typing.Check(
		"isnumber", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}
	value := fmt.Sprintf("%v", args[0])

	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return FALSE
	}
	return TRUE
}
