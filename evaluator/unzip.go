package evaluator

import (
       "mwnci/object"
       "mwnci/typing"
)

func Unzip(args ...object.Object) object.Object {
	if err := typing.Check(
		"unzip", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	Array := args[0].(*object.Array).Elements
	if Array[0].Type() != object.ARRAY_OBJ {
		return newError("argument to unzip() not supported, expected ARRAY[ARRAY], got=ARRAY[%s]", Array[0].Type())
	}
	ABlock := Array[0].(*object.Array).Elements
	DataLen := len(ABlock)
	ZipArray := make([]object.Object, DataLen)
	LenArray := len(Array)
	TempArray := make([]object.Object, LenArray)

	for Index := 0; Index < DataLen; Index++ {
		for In, Block := range Array  {
			if Block.Type() == object.ARRAY_OBJ {
				NBlock := Block.(*object.Array).Elements
				TempArray[In] = NBlock[Index]
				UnMute:=&object.Array{Elements: TempArray}
				ZipArray[Index]=UnMute.Copy()
			} else {
				return newError("argument to unzip() not supported, expected ARRAY[ARRAY], got=ARRAY[%s]", Block.Type())
			}
		}
	}
	return &object.Array{Elements: ZipArray}
}
