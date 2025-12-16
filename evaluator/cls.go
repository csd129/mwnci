package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func Cls(args ...object.Object) object.Object {
	if err := typing.Check(
		"cls", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}

	//
	// To execute "clear", or not
	// There's really no decisive way. So will use
	// escape codes cos this is mainly written for
	// linux at the mo.
	fmt.Print("\033[H\033[2J")
	return NULL
}
