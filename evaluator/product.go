package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"strconv"
)

func Product(args ...object.Object) object.Object {
	if err := typing.Check(
		"product", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) == 0 {
		return &object.Float{Value: 0.0}
	}
	total := float64(1.0)
	for _, el := range arr.Elements {
		num := fmt.Sprintf("%v", el)
		val, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return newError("Unable to parse \"%v\"", el)
		}
		total *= val
	}
	// Fix rounding error
	round := fmt.Sprintf("%.6f", total)
	total, _ = strconv.ParseFloat(round, 64)
	return &object.Float{Value: total}
}
