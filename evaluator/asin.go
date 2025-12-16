package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathAsin(args ...object.Object) object.Object {
	if err := typing.Check(
		"asin", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Asin(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Asin(arg.Value)}
	default:
		return newError("argument to asin() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}
}
