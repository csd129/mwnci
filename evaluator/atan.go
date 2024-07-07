package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathAtan(args ...object.Object) object.Object {
	if err := typing.Check(
		"atan", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Atan(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Atan(arg.Value)}
	default:
		return newError("argument to atan() not supported, expect INTEGER or FLOAT, got %s", args[0].Type())
	}
}
