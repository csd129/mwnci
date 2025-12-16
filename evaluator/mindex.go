package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Mindex(args ...object.Object) object.Object {
	if err := typing.Check(
		"mindex", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError("%s", err.Error())
	}

	haystack := args[0].(*object.Array)
	needle := args[1].(object.Comparable)
	CountArray := make([]object.Object, 0)
	for i := range haystack.Elements {
		if needle.Compare(haystack.Elements[i]) == 0 {
			CountArray=append(CountArray, &object.Integer{Value: int64(i)})
		}
	}
	return &object.Array{Elements: CountArray}
}


