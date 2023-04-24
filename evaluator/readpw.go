package evaluator

import (
    "syscall"
    "mwnci/object"
    "mwnci/typing"
    "golang.org/x/term"
)

func Readpw(args ...object.Object) object.Object {
    if err := typing.Check(
        "readpw", args,
	typing.ExactArgs(0),
    ); err != nil {
        return newError(err.Error())
    }

    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        return newError(err.Error())
    }
    password := string(bytePassword)
    return &object.String{Value: password}
}

func init() {
    RegisterBuiltin("readpw",
        func(env *object.Environment, args ...object.Object) object.Object {
	    return (Readpw(args...))
	})
}