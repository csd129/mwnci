package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrExtend(args ...object.Object) object.Object {
	if err := typing.Check(
		"extend", args,
		typing.MinimumArgs(2),
	); err != nil {
		return newError(err.Error())
	}
	arr1 := args[0].(*object.Array)
	newArray := arr1.Copy()
	for i := 1; i < len(args); i++ {
		if args[i].Type() == object.ARRAY_OBJ {
			for _, v := range args[i].(*object.Array).Elements {
				newArray.Append(v)
			}
		} else {
			return newError("argument to extend() not supported, expected ARRAY, got=%s", args[i].Type())

		}
	}
	return newArray
}

