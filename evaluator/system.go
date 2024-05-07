package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os/exec"
)

// System...
func System(args ...object.Object) object.Object {
	if err := typing.Check(
		"system", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	full_command := args[0].(*object.String).Value

	out, err := exec.Command("bash", "-c", full_command).Output()
	if err != nil {
		return newError("Failed to execute: `%s`: %s", full_command, err.Error())
	}

	return &object.String{Value: string(out)}

}
