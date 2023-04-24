package evaluator

import (
	"fmt"
	"time"
	"mwnci/object"
	"mwnci/typing"
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
	return &object.String{Value: fmt.Sprintf("%s", current_time.Format(time.ANSIC))}
}

func init() {
	RegisterBuiltin("date",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Date(args...))
		})
}
