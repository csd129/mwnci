package evaluator

import (
	"sort"

	"mwnci/object"
	"mwnci/typing"
)

// Sorted ...
func Sortit(args ...object.Object) object.Object {
	if err := typing.Check(
		"sort", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	arr := args[0].(*object.Array)
	if !arr.SameType(arr) {
		return newError("TypeError: Array contents of different types")
	}
	sort.Sort(arr)
	return arr
}
