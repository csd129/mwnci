package evaluator

import (
	"fmt"
	"mwnci/object"
)

// output a string to stdout
func putsFun(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Print(arg.Inspect())
	}
	return &object.Null{}
}

func init() {
	RegisterBuiltin("print",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (putsFun(args...))
		})
}
