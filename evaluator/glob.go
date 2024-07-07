package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"path/filepath"
)

func Glob(args ...object.Object) object.Object {
	if err := typing.Check(
		"glob", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	checkpath := args[0].(*object.String).Value
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

