package evaluator

import (
	"os"

	"mwnci/object"
	"mwnci/typing"
)

// os.setenv( "PATH", "/bin:/usr/bin" );
func setEnvFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"setenv", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	name := args[0].(*object.String).Value
	value := args[1].(*object.String).Value
	err := os.Setenv(name, value)
	if err != nil {
		return FALSE
	} else {
		return TRUE
	}
}
