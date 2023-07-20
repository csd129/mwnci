package evaluator

import (
	b64 "encoding/base64"
	"mwnci/object"
	"mwnci/typing"
)

func B64dec(args ...object.Object) object.Object {
	if err := typing.Check(
		"b64dec", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	data := args[0].Inspect()
	sDec, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return newError(err.Error())
	}
	decoded := string(sDec)
	return &object.String{Value: decoded}
}

func init() {
	RegisterBuiltin("b64dec",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (B64dec(args...))
		})
}
