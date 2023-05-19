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
		return newError(err.Error())
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Cbrt(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Cbrt(v)}
	default:
		return newError("argument to `cbrt` not supported, got=%s",
			args[0].Type())
	}

}

func init() {
	RegisterBuiltin("cbrt",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (mathCbrt(args...))
		})
}
