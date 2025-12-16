package evaluator

import (
	"fmt"
	"os"

	"mwnci/object"
	"mwnci/typing"
)

func Assert(args ...object.Object) object.Object {
	if err := typing.Check(
		"assert", args,
		typing.RangeOfArgs(2, 3),
		typing.WithTypes(object.BOOLEAN_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	if !args[0].(*object.Boolean).Value {
		if len(args) == 3 {
			return FALSE
		}
		fmt.Printf("Assertion Error: %s\n", args[1].(*object.String).Value)
		os.Exit(1)
	}
	return TRUE
}
