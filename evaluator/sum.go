package evaluator

import (
	"strconv"
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)


func Sum(args ...object.Object) object.Object {
	if err := typing.Check(
		"sum", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	total:= float64(0.0)
	for _, el := range arr.Elements {
		num:=fmt.Sprintf("%v", el)
		val, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return newError("Unable to parse \"%v\"", el)
		}
		total += val
	}
	return &object.Float{Value: total}
}

func init() {
	RegisterBuiltin("sum",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Sum(args...))
		})
}
