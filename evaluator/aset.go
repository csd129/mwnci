package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Arrayset(args ...object.Object) object.Object {
	if err := typing.Check(
		"aset", args,
		typing.ExactArgs(3),
	); err != nil {
		return newError(err.Error())
	}
	arr := args[0].(*object.Array)
	nums := arr.Copy()
	elem := int(args[1].(*object.Integer).Value)

	if (elem > len(arr.Elements)-1) || (elem < 0) {
		return newError("IndexError: array index [%d] out of range ", elem)
	} else {
		val := args[2]
		nums.Aset(elem, val)
	}
	return nums

}

func init() {
	RegisterBuiltin("aset",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Arrayset(args...))
		})
}
