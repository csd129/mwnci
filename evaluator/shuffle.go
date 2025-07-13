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
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
//	newa := arr.Copy()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(length, func(i, j int) { arr.Swap(i, j) })
	return arr
}
