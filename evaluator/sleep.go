package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"time"
)

// Sleep ...
func Sleep(args ...object.Object) object.Object {
	if err := typing.Check(
		"sleep", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	timer := int(args[0].(*object.Integer).Value)
	time.Sleep(time.Duration(timer) * time.Millisecond)

	return NULL
}
