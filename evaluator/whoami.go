package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
	"os/user"
)

func Whoami(args ...object.Object) object.Object {
	if err := typing.Check(
		"whoami", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}
	UserName := ""
	currentuser, err := user.Current()
	if err != nil {
		UserName = os.Getenv("USER")
	} else {
		UserName = currentuser.Username
	}
	return &object.String{Value: UserName}
}
