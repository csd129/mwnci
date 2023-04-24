package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"math"
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
		v := arg.Value
		return &object.Float{Value: math.Atan(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Atan(float64(v))}
	default:
		return newError("argument to atan() not supported, expect INTEGER or FLOAT, got %s", args[0].Type())
	}

}

func init() {
	RegisterBuiltin("atan",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathAtan(args...))
		})
}
