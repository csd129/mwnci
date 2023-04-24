package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"io/ioutil"
)

func FCopy(args ...object.Object) object.Object {
	if err := typing.Check(
		"copy", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	orig := args[0].Inspect()
	dest := args[1].Inspect()

	bytesRead, err := ioutil.ReadFile(orig)
	if err != nil {
		return newError(err.Error())
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)
	if err != nil {
		return newError(err.Error())
	}
	return TRUE
}

func init() {
	RegisterBuiltin("copy",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (FCopy(args...))
		})
}
