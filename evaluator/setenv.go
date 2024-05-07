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

	err := os.Setenv(args[0].(*object.String).Value, args[1].(*object.String).Value)
	if err != nil {
		return FALSE
	} else {
		return TRUE
	}
}
