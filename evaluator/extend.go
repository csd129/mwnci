package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrExtend(args ...object.Object) object.Object {
	if err := typing.Check(
		"extend", args,
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	arr1 := args[0].(*object.Array)
	newArray := arr1.Copy()
	for i:=1; i < len(args); i++ {
		for _,v := range args[i].(*object.Array).Elements {
			newArray.Append(v)
		}
	}
	return newArray

}

func init() {
	RegisterBuiltin("extend",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (ArrExtend(args...))
		})
}
