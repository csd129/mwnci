package evaluator

import (
	"strings"
	"mwnci/object"
	"mwnci/typing"
)

func Replace(args ...object.Object) object.Object {
	if err := typing.Check(
		"replace", args,
		typing.ExactArgs(4),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	line := args[0].(*object.String).Value
	search := args[1].(*object.String).Value
	rep := args[2].(*object.String).Value
	count := int(args[3].(*object.Integer).Value)
	out := strings.Replace(line, search, rep, count)
	return &object.String{Value: out}
}

func init() {
	RegisterBuiltin("replace",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Replace(args...))
		})
}
