package evaluator

import (
	"os"

	"mwnci/object"
	"mwnci/typing"
)

// exit a program.
func exitFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"exit", args,
		typing.RangeOfArgs(0, 1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	code := 0

	// Optionally an exit-code might be supplied as an argument
	if len(args) > 0 {
		code = int(args[0].(*object.Integer).Value)
	}

	os.Exit(code)
	// sigh
	return nil
}

