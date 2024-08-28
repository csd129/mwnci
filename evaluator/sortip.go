package evaluator

import (
	"sort"

	"mwnci/object"
	"mwnci/typing"
)

func SortIP(args ...object.Object) object.Object {
	if err := typing.Check(
		"sortips", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr := args[0].(*object.Array)
	newArray := arr.Copy()
	if newArray.SameType(newArray) {
		for i, v := range newArray.Elements {
			if CheckIP(v) == FALSE {
				return newError("Array contains non-IP data [\"%v\"]", v)
			}
			newArray.Aset(i, IP2I(v))
		}
		sort.Sort(newArray)
		for i, v := range newArray.Elements {
			newArray.Aset(i, I2IP(v))
		}
	} else {
		return newError("Array contains mixed data.")
	}
	return newArray
}
