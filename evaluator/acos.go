package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"math"
)

func mathAcos(args ...object.Object) object.Object {
	if err := typing.Check(
		"acos", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Acos(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Acos(v)}
	default:
		return newError("argument to acos() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}

}

func init() {
	RegisterBuiltin("acos",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathAcos(args...))
		})
}
