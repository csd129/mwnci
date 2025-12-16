package evaluator

import (
	"os"
	"strings"
	"mwnci/object"
	"mwnci/typing"
)

func EnvFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"env", args,
		typing.ExactArgs(0),
	); err != nil {
		return newError("%s", err.Error())
	}
	newHash := make(map[object.HashKey]object.HashPair)
	for _, value := range os.Environ() {
		SplitVar:=strings.Split(value, "=")
		key := &object.String{Value: SplitVar[0]}
		val := &object.String{Value: SplitVar[1]}
		newHashPair := object.HashPair{Key: key, Value: val}
		newHash[key.HashKey()] = newHashPair
	}
	return &object.Hash{Pairs: newHash}
}

