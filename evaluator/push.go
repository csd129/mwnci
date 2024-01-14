package evaluator

import (
	"mwnci/object"
)

// push something onto an array
func fpushFun(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got=%s",
			args[0].Type())
	}
	arr := args[0].(*object.Array)
	newArray := arr.Copy()
	newArray.Append(args[1])
	return newArray
}

