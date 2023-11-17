package evaluator

import (
	"fmt"
	"mwnci/object"
)

// output a string to stdout
func Printat(args ...object.Object) object.Object {
	X := int(args[0].(*object.Integer).Value)
	Y := int(args[1].(*object.Integer).Value)
	X = 10
	Y = 10
	Text := args[2].Inspect()
	fmt.Print("\033[", X, ";", Y, "H", Text, "\033[0m")

	return NULL
}

func init() {
	RegisterBuiltin("printat",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Printat(args...))
		})
}
