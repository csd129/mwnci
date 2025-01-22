package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"strconv"
)

// Change a mode of a file - note the second argument is a string
// to emphasise octal.
func chmodFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"chmod", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].String()
	mode := ""

	switch args[1].(type) {
	case *object.String:
		mode = args[1].(*object.String).Value
	default:
		return newError("Second argument must be STRING, got %v", args[1])
	}

	// convert from octal -> decimal
	result, err := strconv.ParseInt(mode, 8, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling chmod(): Unable to create permissions with %v\n", args[1])
		return FALSE
	}

	// Change the mode.
	err = os.Chmod(path, os.FileMode(result))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling chmod(): %v\n", err.Error())
		return FALSE
	}
	return TRUE
}
