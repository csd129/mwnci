package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Get hash items
func hashItems(args ...object.Object) object.Object {
	if err := typing.Check(
		"items", args,
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
	temparray := make([]object.Object, 2)
	// Now copy the keys into it.
	i := 0
	for _, ent := range hash.Pairs {
		temparray[0]=ent.Key
		temparray[1]=ent.Value
		pushlist:=&object.Array{Elements: temparray}
		foo:=pushlist.Copy()
		array[i] = foo
		i++
	}

	// Return the array.
	return &object.Array{Elements: array}
}

