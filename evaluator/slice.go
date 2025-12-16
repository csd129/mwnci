package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Slice(args ...object.Object) object.Object {
	if err := typing.Check(
		"slice", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	Array := args[0].(*object.Array).Elements
	Len := len(Array)
	Start := int(args[1].(*object.Integer).Value)
	End := int(args[2].(*object.Integer).Value)+1
	if Start < 0 || End > Len {
		return newError("Index to slice() is out of range [%d:%d]", Start, End-1)
	}
	ASlice := Array[Start:End]
	return &object.Array{Elements: ASlice}
}
