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
		return newError("%s", err.Error())
	}
	data := []byte(args[0].(*object.String).String())
	return &object.String{Value: base64.StdEncoding.EncodeToString(data)}
}

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
		return newError("%s", err.Error())
	}
	return &object.String{Value: string(decoded)}
}
