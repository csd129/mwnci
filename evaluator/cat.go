package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Cat ...
func ConCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	data := ""
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling cat(): %v\n", err.Error())
		return &object.String{Value: string(data)}
	}
	return &object.String{Value: string(contents)}
}
