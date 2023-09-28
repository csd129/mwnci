package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"os"
	"strings"
)

// Cat ...
func Cat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	data, err := os.ReadFile(filename)
	if err != nil {
		return newError("IOError: %s: %s", filename, err)
	}
	return &object.String{Value: strings.TrimSuffix(string(data), "\n")}
}

func init() {
	RegisterBuiltin("cat",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Cat(args...))
		})
}
