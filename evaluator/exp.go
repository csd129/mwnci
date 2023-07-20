package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)
func mathExp(args ...object.Object) object.Object {
	if err := typing.Check(
		"exp", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Exp(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Exp(arg.Value)}
	default:
		return newError("argument to `exp` not supported, got=%s",
			args[0].Type())
	}
}

func init() {
	RegisterBuiltin("exp",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathExp(args...))
		})
}
