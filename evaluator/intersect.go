package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"fmt"
	"strings"
)

func ArrIntersect(args ...object.Object) object.Object {
	if err := typing.Check(
		"intersect", args,
		typing.MinimumArgs(2),
	); err != nil {
		return newError(err.Error())
	}
	arr1 := args[0].(*object.Array)
	arglen := len(args)
	newArray := arr1.Copy()
	for i := 1; i < arglen; i++ {
		if args[i].Type() == object.ARRAY_OBJ {
			for _, v := range args[i].(*object.Array).Elements {
				newArray.Append(v)
			}
		} else {
			return newError("argument to intersect() not supported, expected ARRAY, got=%s", args[i].Type())

		}
	}
	// All arrays extended into 1
	dict := make(map[string]int)
	buildline := ""
	for _, data := range newArray.Elements {
		HashData := fmt.Sprintf("%v", data)
		dict[HashData]++
	}
	for k, v := range dict {
		if v >= arglen {
			buildline = fmt.Sprintf("%v %v", buildline, k)
		}
	}
	buildline=strings.Trim(buildline, " ")
	tokens := strings.Split(buildline, " ")
	elements := make([]object.Object, len(tokens))
	for i, token := range tokens {
		elements[i] = &object.String{Value: token}
	}
	return &object.Array{Elements: elements}
}

func init() {
	RegisterBuiltin("intersect",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (ArrIntersect(args...))
		})
}
