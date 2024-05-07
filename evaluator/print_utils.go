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
	return NULL
}

// Println ...
func Println(args ...object.Object) object.Object {
	putsFun(args...)
	fmt.Print("\n")
	return NULL
}

// sprintfFun is the implementation of our `sprintf` function.
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

// printfFun is the implementation of our `printf` function.
func printfFun(args ...object.Object) object.Object {

	// Convert to the formatted version, via our `sprintf`
	// function.
	out := sprintfFun(args...)

	// If that returned a string then we can print it
	if out.Type() == object.STRING_OBJ {
		fmt.Print(out.(*object.String).Value)
	}
	return NULL
}
