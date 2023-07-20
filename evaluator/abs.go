package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
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
		return &object.Float{Value: math.Abs(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Abs(arg.Value)}
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
