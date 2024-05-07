package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func Args(args ...object.Object) object.Object {
	if err := typing.Check(
		"args", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	l := len(os.Args[1:])
	result := make([]object.Object, l)
	for i, txt := range os.Args[1:] {
		result[i] = &object.String{Value: txt}
	}
	return &object.Array{Elements: result}
}
