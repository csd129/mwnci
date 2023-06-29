package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os/user"
)

func Whoami(args ...object.Object) object.Object {
	if err := typing.Check(
		"whomai", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	user, err := user.Current()
	if err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: user.Username}
}

func init() {
	RegisterBuiltin("whoami",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Whoami(args...))
		})
}
