package evaluator

import (
	"mwnci/object"
)

// Open a file
func openFun(args ...object.Object) object.Object {

	path := ""
	mode := "r"

	// We need at least one arg
	if len(args) < 1 {
		return newError("wrong number of arguments. got=%d, want=1+",
			len(args))
	}

	// Get the filename
	switch args[0].(type) {
	case *object.String:
		path = args[0].(*object.String).Value
	default:
		return newError("argument to `file` not supported, got=%s",
			args[0].Type())

	}

	// Get the mode (optional)
	if len(args) > 1 {
		switch args[1].(type) {
		case *object.String:
			mode = args[1].(*object.String).Value
		default:
			return newError("argument to `file` not supported, got=%s",
				args[0].Type())

		}
	}

	// Create the object
	file := &object.File{Filename: path}
	file.Open(mode)
	return (file)
}

