package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
	"os/user"
)

// Getegid ...
func Getegid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getegid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	return &object.Integer{Value: int64(os.Getegid())}
}

// Geteuid ...
func Geteuid(args ...object.Object) object.Object {
	if err := typing.Check(
		"geteuid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	return &object.Integer{Value: int64(os.Geteuid())}
}

// Getgid ...
func Getgid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getgid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}

	return &object.Integer{Value: int64(os.Getgid())}
}

// Getuid ...
func Getuid(args ...object.Object) object.Object {
	if err := typing.Check(
		"getuid", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}
	return &object.Integer{Value: int64(os.Getuid())}
}

func Whoami(args ...object.Object) object.Object {
	if err := typing.Check(
		"whoami", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
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
