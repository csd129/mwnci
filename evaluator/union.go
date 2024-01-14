package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrUnion(args ...object.Object) object.Object {
	if err := typing.Check(
		"union", args,
		typing.MinimumArgs(2),
	); err != nil {
		return newError(err.Error())
	}
	newArray := args[0].(*object.Array).Copy()
	for i := 1; i < len(args); i++ {
		if args[i].Type() == object.ARRAY_OBJ {
			for _, v := range args[i].(*object.Array).Elements {
				newArray.Append(v)
			}
		} else {
			return newError("argument to union() not supported, expected ARRAY, got=%s", args[i].Type())

		}
	}
	return Uniq(newArray)
}

