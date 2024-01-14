package evaluator

import (
	"strings"

	"mwnci/object"
	"mwnci/typing"
)

// Fields ...
func Fields(args ...object.Object) object.Object {
	if err := typing.Check(
		"fields", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	s := args[0].(*object.String).Value

	tokens := strings.Fields(s)
	elements := make([]object.Object, len(tokens))
	for i, token := range tokens {
		elements[i] = &object.String{Value: token}
	}
	return &object.Array{Elements: elements}
}

