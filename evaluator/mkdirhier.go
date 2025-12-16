package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"strconv"
)

// Mkdir...
func Mkdirhier(args ...object.Object) object.Object {
	if err := typing.Check(
		"mkdirhier", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	directory := args[0].(*object.String).Value
	perms := args[1].(*object.String).Value
	mode, err := strconv.ParseInt(perms, 8, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling mkdirhier(): Unable to create permissions with %v\n", perms)
		return FALSE
	}
	err = os.MkdirAll(directory, os.FileMode(mode))
	if err != nil && !os.IsExist(err) {
		fmt.Fprintf(os.Stderr, "Error calling mkdirhier(): %v\n", err.Error())
		return FALSE
	}
	return TRUE
}
