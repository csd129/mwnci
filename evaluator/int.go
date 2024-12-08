package evaluator

import (
	"strconv"
	"strings"
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

// convert a double/string to an int
func intFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"int", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}
	switch args[0].(type) {
	case *object.String:
		input := args[0].(*object.String).Value
		// We might have a string that looks like a float,
		// so we only use the numbers to the left of the
		// decimal point
		decimal := strings.Split(input, ".")
		input = decimal[0]
		i, err := strconv.Atoi(input)
		if err == nil {
			return &object.Integer{Value: int64(i)}
		}
		return newError("Converting string '%s' to int failed %s", input, err.Error())

	case *object.Boolean:
		input := args[0].(*object.Boolean).Value
		if input {
			return &object.Integer{Value: 1}
		}
		return &object.Integer{Value: 0}
	case *object.Integer:
		// nop
		return args[0]
	case *object.Float:
		input := args[0].(*object.Float).Value
		return &object.Integer{Value: int64(input)}
	case *object.Array:
		arr := args[0].(*object.Array)
		newArray := arr.Copy()
		for i, v := range newArray.Elements {
			Num:=fmt.Sprintf("%v", v)
			decimal := strings.Split(Num, ".")
			Num=decimal[0]
			foo, err := strconv.Atoi(Num)
			if err == nil {
				newArray.Aset(i, &object.Integer{Value: int64(foo)})
			}
		}
		return newArray
	default:
		return newError("argument to `int` not supported, got=%s",
			args[0].Type())
	}
}

