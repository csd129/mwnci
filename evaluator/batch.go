package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Batch(args ...object.Object) object.Object {
	if err := typing.Check(
		"batch", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	NewArray:=arr.Copy().Elements
	LenArray := len(NewArray)
	limit := args[1].(*object.Integer).Value
	Ilimit := int(limit)
	if Ilimit < 0 {
		return newError("argument to batch() not supported, expected a batch size > 0")
	}
	NewSize:=float64(LenArray)/float64(Ilimit)
	INewSize:=LenArray / Ilimit
	if NewSize != float64(INewSize) { NewSize++ }
	NewList := make([]object.Object, int(NewSize))
	n := 0
	for i := 0; i < LenArray; i += Ilimit {
		batch := NewArray[i:min(i+Ilimit, LenArray)]
		NewList[n]=&object.Array{Elements: batch}
		n++
	}
	return &object.Array{Elements: NewList}
}



func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}
