package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Pop ...
func Pop(args ...object.Object) object.Object {
	if err := typing.Check(
		"pop", args,
		typing.RangeOfArgs(1,2),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	Pos:=int64(0)
	arr := args[0].(*object.Array)
	if len(args) == 2 {
		Pos = args[1].(*object.Integer).Value
	} else {
		Pos = int64(len(arr.Elements)-1)
	}
	if Pos < 0 || Pos > int64(len(arr.Elements)) {
		return newError("IndexError: Index [%d] out of range", Pos)
	}

	element := arr.Elements[Pos]
	harrayDelete(args[0], &object.Integer{Value: Pos})
	return element
}
