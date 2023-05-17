package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrayExtend(args ...object.Object) object.Object {
	if err := typing.Check(
		"extend", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.ARRAY_OBJ, object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	arr1 := args[0].(*object.Array)
	arr2 := args[1].(*object.Array)
	totalcount := len(arr1.Elements) + len(arr2.Elements)
	elements := make([]object.Object, totalcount)
	elements = append(arr1.Elements, arr2.Elements...)
	return &object.Array{Elements: elements}

}

func init() {
	RegisterBuiltin("extend",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (ArrayExtend(args...))
		})
}
