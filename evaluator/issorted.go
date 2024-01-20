package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"sort"
)

func Issorted(args ...object.Object) object.Object {
	if err := typing.Check(
		"issorted", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	if sort.IsSorted(args[0].(*object.Array)) {
		return TRUE
	} else {
		return FALSE
	}
}
