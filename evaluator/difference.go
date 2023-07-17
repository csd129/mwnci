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
			return newError("argument to extend() not supported, expected ARRAY, got=%s", args[i].Type())

		}
	}
	list := make([]object.Object, 0)
	for _, val1 := range arr1.Elements {
		found := false
		value1:=val1.Inspect()
		for _, val2 := range newArray.Elements {
			value2:=val2.Inspect()
			if value1 == value2 {
				found=true
				break
			}
		}
		if !found {
			list = append(list,val1)
		}
	}
	return &object.Array{Elements: list}
}

func init() {
	RegisterBuiltin("difference",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (ArrDiff(args...))
		})
}
