package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"syscall"

	"golang.org/x/term"
)

func Readpw(args ...object.Object) object.Object {
	if err := typing.Check(
		"readpw", args,
		typing.RangeOfArgs(0, 1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	prompt := ""
	if len(args) == 1 {
		prompt = args[0].(*object.String).Value
	} else {
		prompt = "> "
	}
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: string(bytePassword)}
}

