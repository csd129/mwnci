package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathSin(args ...object.Object) object.Object {
	if err := typing.Check(
		"sin", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sin(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sin(arg.Value)}
	default:
		return newError("argument to `sin` not supported, got=%s",
			args[0].Type())
	}
}
