package evaluator

import (
	"strconv"

	"mwnci/object"
	"mwnci/typing"
)

// convert a double/string to a float
func floatFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"float", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	switch args[0].(type) {
	case *object.String:
		input := args[0].(*object.String).Value
		i, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return &object.Float{Value: float64(i)}
		}
		return newError("Converting string '%s' to float failed %s", input, err.Error())

	case *object.Boolean:
		input := args[0].(*object.Boolean).Value
		if input {
			return &object.Float{Value: 1.00}

		}
		return &object.Float{Value: 0.00}
	case *object.Float:
		// nop
		return args[0]
	case *object.Integer:
		input := args[0].(*object.Integer).Value
		return &object.Float{Value: float64(input)}
	default:
		return newError("argument to `float` not supported, got=%s",
			args[0].Type())
	}
}

