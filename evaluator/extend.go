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
	BaseArray := args[0].(*object.Array)
	for i := 1; i < len(args); i++ {
		if args[i].Type() == object.ARRAY_OBJ {
			for _, v := range args[i].(*object.Array).Elements {
				BaseArray.Append(v)
			}
			return BaseArray
		} else if args[i].Type() == object.STRING_OBJ {
			Line := strings.Split(args[i].(*object.String).Value, "")
			for _, v := range Line {
				BaseArray.Append(&object.String{Value: v})
			}
		} else {
			return newError("argument to extend() not supported, expected ARRAY or STRING, got=%s", args[i].Type())

		}
	}
	return BaseArray
}
