package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"fmt"
	"os"
)

func Cd(args ...object.Object) object.Object {
	dir := os.Getenv("HOME")
	if err := typing.Check(
		"cd", args,
		typing.RangeOfArgs(0, 1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	if len(args) == 1 {
		dir = args[0].(*object.String).Value
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling cd(): %v\n", err.Error())
		return FALSE
	}
	return &object.Boolean{Value: true}
}
