package evaluator

import (
	b64 "encoding/base64"
	"mwnci/object"
	"mwnci/typing"
)

func B64urldec(args ...object.Object) object.Object {
	if err := typing.Check(
		"b64urldec", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	data := args[0].Inspect()
	sDec, err := b64.URLEncoding.DecodeString(data)
	if err != nil {
		return newError(err.Error())
	}
	decoded := string(sDec)
	return &object.String{Value: decoded}
}

func init() {
	RegisterBuiltin("b64urldec",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (B64urldec(args...))
		})
}
