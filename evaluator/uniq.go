package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Uniq(args ...object.Object) object.Object {
	if err := typing.Check(
		"uniq", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	arr := args[0].(*object.Array)
	keys := make(map[string]bool)

	list := make([]object.Object, 0)
	for _, entry := range arr.Elements {
		kentry := entry.Inspect()
		if _, value := keys[kentry]; !value {
			keys[kentry] = true
			list = append(list, entry)
		}
	}
	return &object.Array{Elements: list}
}

