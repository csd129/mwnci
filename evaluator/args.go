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
		return newError("%s", err.Error())
	}
	X := false
	l := len(os.Args[1:])
	if l > 0 {
		if os.Args[1] == "-x" {
			X = true
		}
	}
	result := make([]object.Object, l)
	for i, txt := range os.Args[1:] {
		result[i] = &object.String{Value: txt}
	}
	if X == true && l > 1 {
		newresult := Shift(&object.Array{Elements: result})
		return newresult
	}
	return &object.Array{Elements: result}
}
