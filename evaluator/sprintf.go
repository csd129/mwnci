package evaluator

import (
	"fmt"
	"mwnci/object"
)

func sprintfFun(args ...object.Object) object.Object {

	// We expect 1+ arguments
	if len(args) < 1 {
		return NULL
	}

	// Type-check
	if args[0].Type() != object.STRING_OBJ {
		return NULL
	}

	// Get the format-string.
	fs := args[0].(*object.String).Value

	// Convert the arguments to something go's sprintf
	// code will understand.
	argLen := len(args)
	fmtArgs := make([]interface{}, argLen-1)

	// Here we convert and assign.
	for i, v := range args[1:] {
		fmtArgs[i] = v.ToInterface()
	}

	// Call the helper
	out := fmt.Sprintf(fs, fmtArgs...)

	// And now return the value.
	return &object.String{Value: out}
}
