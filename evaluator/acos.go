package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathAcos(args ...object.Object) object.Object {
	if err := typing.Check(
		"acos", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Acos(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Acos(arg.Value)}
	default:
		return newError("argument to acos() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}
}
