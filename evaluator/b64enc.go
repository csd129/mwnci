package evaluator

import (
       b64 "encoding/base64"
       "mwnci/object"
       "mwnci/typing"
)

func B64enc(args ...object.Object) object.Object {
	if err := typing.Check(
		"b64enc", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}     


	data := args[0].Inspect()
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return &object.String{Value: sEnc}
}

func init() {
	RegisterBuiltin("b64enc",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (B64enc(args...))
		})
}
