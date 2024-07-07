package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathCosh(args ...object.Object) object.Object {
	if err := typing.Check(
		"cosh", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cosh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cosh(arg.Value)}
	default:
		return newError("argument to cosh() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}
}

