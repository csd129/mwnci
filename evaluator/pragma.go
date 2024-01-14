package evaluator

import (
	"strings"

	"mwnci/object"
	"mwnci/typing"
)

// set a global pragma
func pragmaFun(args ...object.Object) object.Object {

	// If more than one argument that's an error
	if err := typing.Check(
		"pragma", args,
		typing.RangeOfArgs(0, 1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	// If one argument update to enable the given pragma
	if len(args) == 1 {
		switch args[0].(type) {
		case *object.String:
			input := args[0].(*object.String).Value
			input = strings.ToLower(input)

			if strings.HasPrefix(input, "no-") {
				real := strings.TrimPrefix(input, "no-")
				delete(PRAGMAS, real)
			} else {
				PRAGMAS[input] = 1
			}
		default:
			return newError("argument to `pragma` not supported, got=%s",
				args[0].Type())
		}
	}

	// Now return the pragmas that are in-use.
	len := len(PRAGMAS)

	// Create a new array for the results.
	array := make([]object.Object, len)

	i := 0
	for key := range PRAGMAS {
		array[i] = &object.String{Value: key}
		i++

	}
	return &object.Array{Elements: array}
}

