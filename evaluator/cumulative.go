package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"strconv"
)

func Cumulate(args ...object.Object) object.Object {
	if err := typing.Check(
		"cumulative", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	count_arr := make([]object.Object, len(args[0].(*object.Array).Elements))
	total := float64(0)
	for counter, el := range arr.Elements {
		num := fmt.Sprintf("%v", el)
		val, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return newError("Unable to parse \"%v\" at index [%d]", el, counter)
		}
		total = total + val
		count_arr[counter] = &object.Float{Value: total}
	}
	return &object.Array{Elements: count_arr}
}
