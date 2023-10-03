package evaluator

import (
	"os"
	"os/exec"
	"mwnci/object"
	"mwnci/typing"
)

func Cls(args ...object.Object) object.Object {
	if err := typing.Check(
		"cls", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}     
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return NULL
}

func init() {
	RegisterBuiltin("cls",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Cls(args...))
		})
}
