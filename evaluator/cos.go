package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)
func mathCos(args ...object.Object) object.Object {
	if err := typing.Check(
		"cos", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError( err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Cos(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Cos(float64(v))}
	default:
		return newError("argument to `cos` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}

}

func init() {
	RegisterBuiltin("cos",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathCos(args...))
		})
}
