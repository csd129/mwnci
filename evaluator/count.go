package evaluator

import (
	"strings"
	"mwnci/object"
	"mwnci/typing"
)

func Count(args ...object.Object) object.Object {
	if err := typing.Check(
		"count", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError(err.Error())
	}

	counter := 0
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
		"TypeError: count() expected argument #1 to be `array` got `%s`",args[0].Type(),)
	}
	return &object.Integer{Value: int64(counter)}
}

