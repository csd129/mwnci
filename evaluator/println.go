package evaluator

import (
	"fmt"
	"mwnci/object"
)

// Println ...
func Println(args ...object.Object) object.Object {
	for _, s := range args {
		fmt.Print(s.Inspect())
	}
	fmt.Print("\n")
	return nil
}

func init() {
	RegisterBuiltin("println",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Println(args...))
		})
}
