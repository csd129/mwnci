package evaluator

import (
	"strings"
	"mwnci/object"
	"mwnci/typing"
)

func Replaceall(args ...object.Object) object.Object {
	if err := typing.Check(
		"replaceall", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	line := args[0].(*object.String).Value
	search := args[1].(*object.String).Value
	rep := args[2].(*object.String).Value
	out := strings.ReplaceAll(line, search, rep)
	return &object.String{Value: out}
}

func init() {
	RegisterBuiltin("replaceall",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Replaceall(args...))
		})
}
