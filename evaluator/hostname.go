package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Hostname ...
func Hostname(args ...object.Object) object.Object {
	if err := typing.Check(
		"hostname", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}
	hostname, err := os.Hostname()
	if err != nil {
		return &object.String{Value: "localhost"}
	}
	return &object.String{Value: hostname}
}
