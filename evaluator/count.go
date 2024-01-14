package evaluator

import (
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
	if haystack, ok := args[0].(*object.Array); ok {
		needle := args[1].(object.Comparable)
		for i := range haystack.Elements {
			if needle.Compare(haystack.Elements[i]) == 0 {
				counter++
			}
		}
	} else {
		return newError(
			"TypeError: count() expected argument #1 to be `array` got `%s`",args[0].Type(),)
	}

	return &object.Integer{Value: int64(counter)}
}

