package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Delete a given hash-key
func hashDelete(args ...object.Object) object.Object {
     if err := typing.Check(
		"delete", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError(err.Error())
	} 
	if args[0].Type() != object.HASH_OBJ {
		return newError("argument to delete() must be HASH, got=%s", args[0].Type())
	}

	// The object we're working with
	hash := args[0].(*object.Hash)

	// The key we're going to delete
	key, ok := args[1].(object.Hashable)
	if !ok {
		return newError("key delete() into HASH must be Hashable, got=%s",args[1].Type())
	}

	// Make a new hash
	newHash := make(map[object.HashKey]object.HashPair)

	// Copy the values EXCEPT the one we have.
	for k, v := range hash.Pairs {
		if k != key.HashKey() {
			newHash[k] = v
		}
	}
	return &object.Hash{Pairs: newHash}
}

func init() {
	RegisterBuiltin("delete",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (hashDelete(args...))
		})
}
