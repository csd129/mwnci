package evaluator

import (
	"fmt"
	"mwnci/object"
)

// Println ...
func Println(args ...object.Object) object.Object {
	putsFun(args...)
	fmt.Print("\n")
	return NULL
}

func init() {
	RegisterBuiltin("println",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Println(args...))
		})
}
