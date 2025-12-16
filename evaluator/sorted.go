
package evaluator

import (
	"sort"

	"mwnci/object"
	"mwnci/typing"
)

// Sorted ...
func Sorted(args ...object.Object) object.Object {
	if err := typing.Check(
		"sort", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	arr := args[0].(*object.Array)
	newArray := arr.Copy()
	sort.Sort(newArray)
	return newArray
}

