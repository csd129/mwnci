package evaluator

import (
	"os"

	"mwnci/object"
	"mwnci/typing"
)

// os.getenv( "PATH" ) -> string
func getEnvFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"getenv", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	input := args[0].(*object.String).Value
	return &object.String{Value: os.Getenv(input)}

}

