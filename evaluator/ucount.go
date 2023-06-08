package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func Ucount(args ...object.Object) object.Object {
	if err := typing.Check(
		"ucount", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	dict := make(map[string]int)
	arr := args[0].(*object.Array)
	for _,data := range arr.Elements{
		HashData := fmt.Sprintf("%v", data)
		dict[HashData] = dict[HashData] + 1
	}
	return FALSE
}

func init() {
	RegisterBuiltin("yewcount",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Ucount(args...))
		})
}
