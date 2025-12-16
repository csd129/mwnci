package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func unlinkFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"unlink", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	path := args[0].String()

	err := os.Remove(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling unlink(): %v\n", err.Error())
		return FALSE
	}
	return TRUE
}
