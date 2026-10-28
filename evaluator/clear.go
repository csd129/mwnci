package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Clear(args ...object.Object) object.Object {
	if err := typing.Check(
		"clear", args,
		typing.ExactArgs(1),
		//		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	if args[0].Type() == object.ARRAY_OBJ {
		arr := args[0].(*object.Array)
		arr.Elements = arr.Elements[:0]
		return &object.Array{Elements: arr.Elements}
	}
	if args[0].Type() == object.HASH_OBJ {
		hash := args[0].(*object.Hash)
		for _, ent := range hash.Pairs {
			harrayDelete(hash, ent.Key)
		}
		return &object.Hash{Pairs: hash.Pairs}
	}
	return newError("argument to `clear` not supported, expected ARRAY or HASH, got=%s", args[0].Type())
}
