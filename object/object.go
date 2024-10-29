
// Package object contains our core-definitions for objects.
package object

import (
	"fmt"
)
// Type describes the type of an object.
type Type string

// pre-defined constant Type
const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	FILE_OBJ         = "FILE"
	REGEXP_OBJ       = "REGEXP"
)
// Comparable is the interface for comparing two Object and their underlying
// values. It is the responsibility of the caller (left) to check for types.
// Returns `true` iif the types and values are identical, `false` otherwise.
type Comparable interface {
        Compare(other Object) int
}

// Object is the interface that all of our various object-types must implement.
type Object interface {
	fmt.Stringer

	// Type returns the type of this object.
	Type() Type
	//	Bool() bool
	//	Compare(Object) int
	// Inspect returns a string-representation of the given object.
	Inspect() string

	// InvokeMethod invokes a method against the object.
	// (Built-in methods only.)
	InvokeMethod(method string, env Environment, args ...Object) Object

	// ToInterface converts the given object to a "native" golang value,
	// which is required to ensure that we can use the object in our
	// `sprintf` or `printf` primitives.
	ToInterface() interface{}
}

// Hashable type can be hashed
type Hashable interface {

	// HashKey returns a hash key for the given object.
	HashKey() HashKey
}

// Iterable is an interface that some objects might support.
//
// If this interface is implemented then it will be possible to
// use the `foreach` function to iterate over the object.  If
// the interface is not implemented then a run-time error will
// be generated instead.
type Iterable interface {

	// Reset the state of any previous iteration.
	Reset()

	// Get the next "thing" from the object being iterated
	// over.
	//
	// The return values are the item which is to be returned
	// next, the index of that object, and finally a boolean
	// to say whether the function succeeded.
	//
	// If the boolean value returned is false then that
	// means the iteration has completed and no further
	// items are available.
	Next() (Object, Object, bool)
}

// Sizeable is the interface for returning the size of an Object.
// Object(s) that have a valid size must implement  this interface and the
// Len() method.
type Sizeable interface {
        Len() int
}
