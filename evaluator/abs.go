package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Abs(args ...object.Object) object.Object {
	if err := typing.Check(
		"abs", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch args[0].(type) {
	case *object.Integer:
		value := args[0].(*object.Integer).Value
		if value < 0 {
			value=value * -1
		}
		return &object.Integer{Value: value}
	case *object.Float:
		value := args[0].(*object.Float).Value
		if value < 0 {
			value=value * -1
		}
		return &object.Float{Value: value}
	default:
		return newError("argument to `abs` not supported, expected INTEGER or FLOAT, got %s", args[0].Type())
	}
}
