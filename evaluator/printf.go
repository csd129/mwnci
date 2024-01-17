package evaluator

import (
	"fmt"
	"mwnci/object"
)

// printfFun is the implementation of our `printf` function.
func printfFun(args ...object.Object) object.Object {

	// Convert to the formatted version, via our `sprintf`
	// function.
	out := sprintfFun(args...)

	// If that returned a string then we can print it
	if out.Type() == object.STRING_OBJ {
		fmt.Print(out.(*object.String).Value)
	}
	return &object.Null{}
}
