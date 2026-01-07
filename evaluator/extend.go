package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func ArrExtend(args ...object.Object) object.Object {
	if err := typing.Check(
		"extend", args,
		typing.MinimumArgs(2),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	arr1 := args[0].(*object.Array)
	newArray := arr1.Copy()
	for i := 1; i < len(args); i++ {
		if args[i].Type() == object.ARRAY_OBJ {
			for _, v := range args[i].(*object.Array).Elements {
				newArray.Append(v)
			}
		} else if args[i].Type() == object.STRING_OBJ {
			Line := strings.Split(args[i].(*object.String).Value, "")
			for _, v := range Line {
				newArray.Append(&object.String{Value: v})
			}
		} else {
			return newError("argument to extend() not supported, expected ARRAY or STRING, got=%s", args[i].Type())

		}
	}
	return newArray
}
