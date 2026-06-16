package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Index ...
func Bsearch(args ...object.Object) object.Object {
	if err := typing.Check(
		"bsearch", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError("%s", err.Error())
	}

	if haystack, ok := args[0].(*object.Array); ok {
		needle := args[1].(object.Comparable)
		low := 0
		high := len(haystack.Elements)-1
		for low <= high {
			mid := (low + high)/2
			if needle.Compare(haystack.Elements[mid]) == 0 {
				return &object.Integer{Value: int64(mid)}
			}
			if needle.Compare(haystack.Elements[mid]) == -1 {
				high = mid-1
			} else {
				low = mid+1
			}
		}
		return &object.Integer{Value: int64(-1)}
	}

	return newError(
		"TypeError: bsearch() expected argument #1 to be `array` got `%s`", args[0].Type(),
	)
}
