package evaluator

import (
       b64 "encoding/base64"
       "fmt"
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
	sDec, _ := b64.StdEncoding.DecodeString(data)
	decoded := fmt.Sprintf("%s", sDec)
	return &object.String{Value: decoded}
}

func init() {
	RegisterBuiltin("b64dec",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (B64dec(args...))
		})
}
