package evaluator

import (
	"strings"
	"mwnci/object"
	"mwnci/typing"
)

// Join ...
func Joiner(args ...object.Object) object.Object {
	if err := typing.Check(
		"join", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.ARRAY_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	sep := args[1].(*object.String)
	a := make([]string, len(arr.Elements))
	for i, el := range arr.Elements {
		a[i] = el.Inspect()
	}
	return &object.String{Value: strings.Join(a, sep.Value)}
}




