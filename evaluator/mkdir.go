package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"strconv"
)

func mkdirFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"mkdir", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	path := args[0].(*object.String).Value
	mode, err := strconv.ParseInt("755", 8, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling chmod(): Unable to create permissions with 755\n")
		return FALSE
	}
	err = os.Mkdir(path, os.FileMode(mode))
	if err != nil && !os.IsExist(err) {
		fmt.Fprintf(os.Stderr, "Error calling mkdir(): %v\n", err.Error())
		return FALSE
	} else {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling mkdir(): %v\n", err.Error())
			return FALSE
		}
	}
	return TRUE
}
