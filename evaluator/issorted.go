package evaluator

import (
	"sort"
	"mwnci/object"
	"mwnci/typing"
)

func Issorted(args ...object.Object) object.Object {
	if err := typing.Check(
		"issorted", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}
	
	foo := args[0].(*object.Array)
	nums := foo.Copy()
        if sort.IsSorted(nums) {
		return TRUE
	} else {
		return FALSE
	}
}

func init() {
	RegisterBuiltin("issorted",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Issorted(args...))
		})
}
