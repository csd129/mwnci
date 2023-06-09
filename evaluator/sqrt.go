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
		return newError(err.Error())
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Sqrt(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Sqrt(v)}
	default:
		return newError("argument to `sqrt` not supported, got=%s",
			args[0].Type())
	}

}

func init() {
	RegisterBuiltin("sqrt",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathSqrt(args...))
		})
}
