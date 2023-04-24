package evaluator

import (
	"os"
	"mwnci/object"
	"mwnci/typing"
)

// Getgroups ...
func GetGroups(args ...object.Object) object.Object {
	if err := typing.Check(
		"getgroups", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError(err.Error())
	}

	groups,_ := os.Getgroups()
	elements := make([]object.Object, len(groups))
	for i, num := range groups {
		elements[i] = &object.Integer{Value: int64(num)}
	}
	return &object.Array{Elements: elements}
}

func init() {
	RegisterBuiltin("getgroups",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (GetGroups(args...))
		})
}
