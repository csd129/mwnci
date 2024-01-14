package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"time"
)

// Date...
func Date(args ...object.Object) object.Object {
	if err := typing.Check(
		"date", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}

	current_time := time.Now()
	return &object.String{Value: current_time.Format(time.ANSIC)}
}

