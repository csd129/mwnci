package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrayInsert(args ...object.Object) object.Object {
	if err := typing.Check(
		"insert", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	arr := args[0].(*object.Array)
//	newArray := arr.Copy()
	elem := int(args[1].(*object.Integer).Value)

	if (elem > len(arr.Elements)-1) || (elem < 0) {
		return newError("IndexError: array index [%d] out of range ", elem)
	} else {
		val := args[2]
		arr.Insert(elem, val)
	}
	return arr

}

