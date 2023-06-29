package evaluator

import (
	b64 "encoding/base64"
	"mwnci/object"
	"mwnci/typing"
)

func B64urlenc(args ...object.Object) object.Object {
	if err := typing.Check(
		"b64urlenc", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	data := args[0].Inspect()
	sEnc := b64.URLEncoding.EncodeToString([]byte(data))
	return &object.String{Value: sEnc}
}

func init() {
	RegisterBuiltin("b64urlenc",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (B64urlenc(args...))
		})
}
