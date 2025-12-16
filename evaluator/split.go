package evaluator

import (
	"strings"

	"mwnci/object"
	"mwnci/typing"
)

// Split ...
func Split(args ...object.Object) object.Object {
	if err := typing.Check(
		"split", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	var sep string
	s := args[0].(*object.String).Value

	if len(args) == 2 {
		sep = args[1].(*object.String).Value
	}

	if len(s) == 0 {
		return NULL
	}
	tokens := strings.Split(s, sep)
	elements := make([]object.Object, len(tokens))
	for i, token := range tokens {
		elements[i] = &object.String{Value: token}
	}
	return &object.Array{Elements: elements}
}
