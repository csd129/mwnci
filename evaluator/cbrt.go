package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathCbrt(args ...object.Object) object.Object {
	if err := typing.Check(
		"cbrt", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError( err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cbrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cbrt(arg.Value)}
	default:
		return newError("argument to `cbrt` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}
}
