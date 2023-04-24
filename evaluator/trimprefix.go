package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Trimprefix(args ...object.Object) object.Object {
	if err := typing.Check(
		"trimprefix", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	line := args[0].(*object.String).Value
	prefix := args[1].(*object.String).Value
	s := strings.TrimPrefix(line, prefix)
	return &object.String{Value: s}
}

func init() {
	RegisterBuiltin("trimprefix",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Trimprefix(args...))
		})
}
