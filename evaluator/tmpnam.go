package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Tmpnam...
func Tmpnam(args ...object.Object) object.Object {
	if err := typing.Check(
		"tmpnam", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	dir := args[0].(*object.String).Value
	prefix := ""
	if len(args) == 2 {
		prefix = args[1].(*object.String).Value
	}
	file, err := os.CreateTemp(dir, prefix)
	if err != nil {
		return newError(err.Error())
	}

	return &object.String{Value: file.Name()}
}
