package evaluator

import (
	"encoding/base64"
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func Base64Dec(args ...object.Object) object.Object {
	if err := typing.Check(
		"base64dec", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	data := args[0].(*object.String).String()
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling base64dec(): %v\n", err.Error())
		return NULL
	}
	return &object.String{Value: string(decoded)}
}
