package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrayRemove(args ...object.Object) object.Object {
	if err := typing.Check(
		"remove", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError(err.Error())
	}
	arr := args[0].(*object.Array)
	nums := arr.Copy()
	elem := int(args[1].(*object.Integer).Value)

	if (elem > len(arr.Elements)-1) || (elem < 0) {
		return newError("IndexError: array index [%d] out of range ", elem)
	}
	copy(nums.Elements[elem:], nums.Elements[elem+1:])
	len := len(nums.Elements)
	nums.Elements = nums.Elements[:len-1]
	return nums
}

func init() {
	RegisterBuiltin("remove",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (ArrayRemove(args...))
		})
}
