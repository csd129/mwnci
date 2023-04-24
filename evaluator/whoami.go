package evaluator

import (
	"os/user"
	"mwnci/object"
	"mwnci/typing"
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
	username := user.Username
	return &object.String{Value: username}
}

func init() {
	RegisterBuiltin("whoami",
		func(env *object.Environment, args ...object.Object) object.Object {
			return(Whoami(args...))
		})
}
