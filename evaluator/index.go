package evaluator

import (
	"strings"

	"mwnci/object"
	"mwnci/typing"
)

// Index ...
func Index(args ...object.Object) object.Object {
	if err := typing.Check(
		"index", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError(err.Error())
	}

	// index("foobar", "bo")
	if haystack, ok := args[0].(*object.String); ok {
		if err := typing.Check(
			"index", args,
			typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
		); err != nil {
			return newError(err.Error())
		}

		needle := args[1].(*object.String)
		index := strings.Index(haystack.Value, needle.Value)
		return &object.Integer{Value: int64(index)}
	}

	// index([1, 2, 3], 2)
	if haystack, ok := args[0].(*object.Array); ok {
		needle := args[1].(object.Comparable)
		for i := range haystack.Elements {
			if needle.Compare(haystack.Elements[i]) == 0 {
				return &object.Integer{Value: int64(i)}
			}
		}
		return &object.Integer{Value: -1}
	}

	return newError(
		"TypeError: index() expected argument #1 to be `array` or `str` got `%s`",
		args[0].Type(),
	)
}

