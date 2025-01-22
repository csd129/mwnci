package evaluator

import (
	"encoding/base64"
	"mwnci/object"
	"mwnci/typing"
)

func Base64Enc(args ...object.Object) object.Object {
	if err := typing.Check(
		"base64enc", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	data := []byte(args[0].(*object.String).String())
	return &object.String{Value: base64.StdEncoding.EncodeToString(data)}
}
