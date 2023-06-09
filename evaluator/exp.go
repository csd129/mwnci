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

	switch args[0].(type) {
	case *object.Integer:
		input := args[0].(*object.Integer).Value
		result := math.Exp(float64(input))
		return &object.Float{Value: float64(result)}
	case *object.Float:
		input := args[0].(*object.Float).Value
		result := math.Exp(input)
		return &object.Float{Value: float64(result)}
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
