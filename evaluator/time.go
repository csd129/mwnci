package evaluator

import (
	"time"
	"mwnci/object"
	"mwnci/typing"
)

// Time...
func Time(args ...object.Object) object.Object {
	if err := typing.Check(
		"time", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	now := time.Now()
	return &object.Integer{Value: now.Unix()}
}

