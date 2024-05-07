package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"encoding/base64"
)

func Base64Enc(args ...object.Object) object.Object {
	if err := typing.Check(
		"base64enc", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	data:=[]byte(args[0].(*object.String).Inspect())
	return &object.String{Value: base64.StdEncoding.EncodeToString(data)}
}
