package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Get hash keys
func hashKeys(args ...object.Object) object.Object {
	if err := typing.Check(
		"keys", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.HASH_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	// The object we're working with
	hash := args[0].(*object.Hash)
	ents := len(hash.Pairs)

	// Create a new array for the results.
	array := make([]object.Object, ents)

	// Now copy the keys into it.
	i := 0
	for _, ent := range hash.Pairs {
		array[i] = ent.Key
		i++
	}

	// Return the array.
	return &object.Array{Elements: array}
}

