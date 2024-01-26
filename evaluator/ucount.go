package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func Ucount(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check(
		"ucount", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	newHash := make(map[object.HashKey]object.HashPair)
	arr := args[0].(*object.Array)
	dict := make(map[string]int)
	for _, data := range arr.Elements {
		HashData := fmt.Sprintf("%v", data)
		dict[HashData]++
	}

	for k, v := range dict {
		key := &object.String{Value: k}
		val := &object.Integer{Value: int64(v)}
		newHashPair := object.HashPair{Key: key, Value: val}
		newHash[key.HashKey()] = newHashPair
	}
	return &object.Hash{Pairs: newHash}
}
