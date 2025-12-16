package evaluator

import (
	"os"

	"mwnci/object"
	"mwnci/typing"
)

// os.setenv( "PATH", "/bin:/usr/bin" );
func unsetEnvFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"unsetenv", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	err := os.Unsetenv(args[0].(*object.String).Value)
	if err != nil {
		return FALSE
	} else {
		return TRUE
	}
}
