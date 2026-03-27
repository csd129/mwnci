package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func acopyFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"copy", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}
	if args[0].Type() == object.ARRAY_OBJ {
		Var := args[0].(*object.Array)
		newVar := Var.Copy()
		return newVar
	}
	if args[0].Type() == object.HASH_OBJ {
		Var := args[0].(*object.Hash)
		newVar := Var.Copy()
		return newVar
	}
	return newError("argument to copy() not supported, expected HASH or ARRAY, got=%s", args[0].Type())
}
