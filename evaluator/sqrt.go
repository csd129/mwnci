package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathSqrt(args ...object.Object) object.Object {
	if err := typing.Check(
		"sqrt", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError( err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sqrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sqrt(arg.Value)}
	default:
		return newError("argument to `sqrt` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}
}
