package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathLogE(args ...object.Object) object.Object {
	if err := typing.Check(
		"loge", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Log(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Log(float64(v))}
	default:
		return newError("argument to `loge` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}
}
