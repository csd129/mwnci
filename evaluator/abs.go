package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"math"
)

// Abs ...
func Abs(args ...object.Object) object.Object {
	if err := typing.Check(
		"abs", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Abs(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Abs(float64(v))}
	default:
		return newError("argument to abs() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}
}

func init() {
	RegisterBuiltin("abs",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Abs(args...))
		})
}
