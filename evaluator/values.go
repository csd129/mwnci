package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Get hash values
func hashValues(args ...object.Object) object.Object {
	if err := typing.Check(
		"values", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.HASH_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	// The object we're working with
	hash := args[0].(*object.Hash)
	ents := len(hash.Pairs)

	// Create a new array for the results.
	array := make([]object.Object, ents)

	// Now copy the keys into it.
	i := 0
	for _, ent := range hash.Pairs {
		array[i] = ent.Value
		i++
	}

	// Return the array.
	return &object.Array{Elements: array}
}

