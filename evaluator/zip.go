package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func zipFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"zip", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.ARRAY_OBJ,object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	arr1 := args[0].(*object.Array)
	arr2 := args[1].(*object.Array)
	lenarr1 := len(arr1.Elements)-1
	lenarr2 := len(arr2.Elements)-1
	total := 0
	if lenarr1 > lenarr2 {
		total=lenarr2
	} else {
		total=lenarr1
	}

	counter := 0
	newHash := make(map[object.HashKey]object.HashPair)
	for counter <= total {
		K:=fmt.Sprintf("%v", arr1.Elements[counter])
		k:=&object.String{Value: K}
		v := arr2.Elements[counter]
		newHashPair := object.HashPair{Key: k, Value: v}
		newHash[k.HashKey()] = newHashPair
		counter++
	}
	return &object.Hash{Pairs: newHash}
}

func init() {
	RegisterBuiltin("zip",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (zipFun(args...))
		})
}
