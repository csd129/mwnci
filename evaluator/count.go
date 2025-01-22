package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Count(args ...object.Object) object.Object {
	if err := typing.Check(
		"count", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError(err.Error())
	}

	counter := 0
	if args[1].Type() == object.HASH_OBJ || args[1].Type() == object.ARRAY_OBJ {
		return newError("TypeError: count() argument #2 cannot be ARRAY or HASH")
	}
	switch args[0].(type) {
	case *object.String:
		needle := args[1].(*object.String).Value
		haystack := args[0].(*object.String).Value
		counter = strings.Count(haystack, needle)
	case *object.Array:
		needle := args[1].(object.Comparable)
		haystack := args[0].(*object.Array)
		for i := range haystack.Elements {
			if needle.Compare(haystack.Elements[i]) == 0 {
				counter++
			}
		}
	default:
		return newError(
			"TypeError: count() expected argument #1 to be `array` or `string` got `%s`", args[0].Type())
	}
	return &object.Integer{Value: int64(counter)}
}
