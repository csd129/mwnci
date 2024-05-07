package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func mathAcos(args ...object.Object) object.Object {
	if err := typing.Check(
		"acos", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Acos(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Acos(arg.Value)}
	default:
		return newError("argument to acos() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}

}

func mathAsin(args ...object.Object) object.Object {
	if err := typing.Check(
		"asin", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Asin(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Asin(arg.Value)}
	default:
		return newError("argument to asin() not supported, expect INTEGER or FLOAT, got %s", args[0].Type())
	}

}

func mathAtan(args ...object.Object) object.Object {
	if err := typing.Check(
		"atan", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Atan(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Atan(arg.Value)}
	default:
		return newError("argument to atan() not supported, expect INTEGER or FLOAT, got %s", args[0].Type())
	}

}

func mathCos(args ...object.Object) object.Object {
	if err := typing.Check(
		"cos", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError( err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cos(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cos(arg.Value)}
	default:
		return newError("argument to `cos` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}

}

func mathCosh(args ...object.Object) object.Object {
	if err := typing.Check(
		"cosh", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cosh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cosh(arg.Value)}
	default:
		return newError("argument to cosh() not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}

}

func mathLog10(args ...object.Object) object.Object {
	if err := typing.Check(
		"log10", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError( err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		return &object.Float{Value: math.Log10(float64(v))}
	case *object.Float:
		v := arg.Value
		return &object.Float{Value: math.Log10(float64(v))}
	default:
		return newError("argument to `log10` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}

}

func mathLogE(args ...object.Object) object.Object {
	if err := typing.Check(
		"loge", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
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

func mathSin(args ...object.Object) object.Object {
	if err := typing.Check(
		"sin", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sin(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sin(arg.Value)}
	default:
		return newError("argument to `sin` not supported, got=%s",
			args[0].Type())
	}

}

func mathSinh(args ...object.Object) object.Object {
	if err := typing.Check(
		"sinh", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sinh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sinh(arg.Value)}
	default:
		return newError("argument to `sinh` not supported, got=%s",
			args[0].Type())
	}

}

func mathTan(args ...object.Object) object.Object {
	if err := typing.Check(
		"tan", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Tan(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Tan(arg.Value)}
	default:
		return newError("argument to `tan` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}

}

func mathTanh(args ...object.Object) object.Object {
	if err := typing.Check(
		"tanh", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Tanh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Tanh(arg.Value)}
	default:
		return newError("argument to `tanh` not supported, expected INTEGER or FLOAT, got=%s", args[0].Type())
	}

}

