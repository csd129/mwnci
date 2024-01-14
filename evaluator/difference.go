package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrDiff(args ...object.Object) object.Object {
	if err := typing.Check(
		"difference", args,
		typing.MinimumArgs(2),
	); err != nil {
		return newError(err.Error())
	}
	arr1 := args[0].(*object.Array)
	newArray := args[1].(*object.Array).Copy()
	for i := 2; i < len(args); i++ {
		if args[i].Type() == object.ARRAY_OBJ {
			for _, v := range args[i].(*object.Array).Elements {
				newArray.Append(v)
			}
		} else {
			return newError("argument to difference() not supported, expected ARRAY, got=%s", args[i].Type())

		}
	}
	list := make([]object.Object, 0)
	for _, val1 := range arr1.Elements {
		found := false
		for _, val2 := range newArray.Elements {
			if val1.Inspect() == val2.Inspect() {
				found = true
				break
			}
		}
		if !found {
			list = append(list, val1)
		}
	}
	return &object.Array{Elements: list}
}

