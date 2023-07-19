package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)
func mathTanh(args ...object.Object) object.Object {
	if err := typing.Check(
		"tanh", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError( err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Tanh(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Tanh(float64(v))}
	default:
		return newError("argument to `tanh` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}

}

func init() {
	RegisterBuiltin("tanh",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathTanh(args...))
		})
}
