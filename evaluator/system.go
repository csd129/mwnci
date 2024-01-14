package evaluator

import (
	"os/exec"
	"mwnci/object"
	"mwnci/typing"
)

// System...
func System(args ...object.Object) object.Object {
	if err := typing.Check(
		"system", args,
		typing.RangeOfArgs(1,2),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	
	full_command := args[0].(*object.String).Value

	out, err := exec.Command("bash", "-c", full_command).Output()
	if err != nil {
	    return newError(err.Error())
	}
        if len(out) > 1 {
		out = out[:len(out)-1]
	}
	return &object.String{Value: string(out)}

}

