package evaluator

import (
	"fmt"
	"path/filepath"
	"mwnci/object"
	"mwnci/typing"
)

func Glob(args ...object.Object) object.Object {
	if err := typing.Check(
		"glob", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	checkpath := fmt.Sprintf("%s",args[0].(*object.String))
	files, err := filepath.Glob(checkpath)
	if err != nil {
		return newError(err.Error())
	}
	elements := make([]object.Object, len(files))

	for i, token := range files {
		elements[i] = &object.String{Value: token}
	}
	return &object.Array{Elements: elements}
}

func init() {
	RegisterBuiltin("glob",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Glob(args...))
		})
}
