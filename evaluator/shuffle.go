package evaluator

import (
	"math/rand"
	"mwnci/object"
	"mwnci/typing"
	"time"
)

func ShuffleFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"shuffle", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	foo := args[0].(*object.Array)
	length := len(foo.Elements)
	newa := foo.Copy()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(length, func(i, j int) { newa.Swap(i, j) })
	return newa
}

func init() {
	RegisterBuiltin("shuffle",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (ShuffleFun(args...))
		})
}
