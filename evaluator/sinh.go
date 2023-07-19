package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"math"
)

func mathSinh(args ...object.Object) object.Object {
	if err := typing.Check(
		"sinh", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Sinh(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Sinh(v)}
	default:
		return newError("argument to `sinh` not supported, got=%s",
			args[0].Type())
	}

}

func init() {
	RegisterBuiltin("sinh",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathSinh(args...))
		})
}
