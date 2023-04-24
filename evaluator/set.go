package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// set a hash-field
func setFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"set", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.HASH_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	key, ok := args[1].(object.Hashable)
	if !ok {
		return newError("key `set` into HASH must be Hashable, got=%s",
			args[1].Type())
	}
	newHash := make(map[object.HashKey]object.HashPair)
	hash := args[0].(*object.Hash)
	for k, v := range hash.Pairs {
		newHash[k] = v
	}
	newHashKey := key.HashKey()
	newHashPair := object.HashPair{Key: args[1], Value: args[2]}
	newHash[newHashKey] = newHashPair
	return &object.Hash{Pairs: newHash}
}

func init() {
	RegisterBuiltin("set",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (setFun(args...))
		})
}
