package evaluator

import (
	"os"
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"encoding/base64"
)

func Base64Dec(args ...object.Object) object.Object {
	if err := typing.Check(
		"base64dec", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	data := args[0].(*object.String).Inspect()
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling base64dec(): %v\n", err.Error())
		return NULL
	}
	return &object.String{Value: string(decoded)}
}

