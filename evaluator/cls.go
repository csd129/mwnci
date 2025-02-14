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
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
