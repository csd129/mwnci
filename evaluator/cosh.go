package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"math"
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
		v := arg.Value
		return &object.Float{Value: math.Cosh(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Cosh(v)}
	default:
		return newError("argument to cosh() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}

}

func init() {
	RegisterBuiltin("cosh",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathCosh(args...))
		})
}
