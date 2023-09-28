package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func Ceil(args ...object.Object) object.Object {
	if err := typing.Check(
		"ceil", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: int64(args[0].(*object.Integer).Value)}
	case *object.Float:
		return &object.Integer{Value: int64(math.Ceil(args[0].(*object.Float).Value))}
	default:
		return newError("argument to ceil() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}
}

func init() {
	RegisterBuiltin("ceil",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Ceil(args...))
		})
}
