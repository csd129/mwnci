package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
	"os/user"
)

func Whoami(args ...object.Object) object.Object {
	if err := typing.Check(
		"whomai", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	UserName := ""
	currentuser, err := user.Current()
	UserName = currentuser.Username
	if err != nil {
		UserName = os.Getenv("USER")
	}
	return &object.String{Value: UserName}
}

func init() {
	RegisterBuiltin("whoami",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Whoami(args...))
		})
}
