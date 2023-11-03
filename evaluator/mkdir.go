package evaluator

import (
	"os"
	"strconv"

	"mwnci/object"
	"mwnci/typing"
)

// mkdir
func mkdirFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"mkdir", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].(*object.String).Value
	mode, err := strconv.ParseInt("755", 8, 64)
	if err != nil {
		return FALSE
	}

	err = os.Mkdir(path, os.FileMode(mode))
	if err != nil && !os.IsExist(err) {
		return FALSE
	}
	return TRUE
}

func init() {
	RegisterBuiltin("mkdir",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mkdirFun(args...))
		})
}
