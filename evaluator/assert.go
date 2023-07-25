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
		typing.ExactArgs(2),
		typing.WithTypes(object.BOOLEAN_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	if !args[0].(*object.Boolean).Value {
		fmt.Printf("Assertion Error: %s", args[1].(*object.String).Value)
		os.Exit(1)
	}
	return nil
}
	
func init() {
	RegisterBuiltin("assert",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Assert(args...))
		})
}
