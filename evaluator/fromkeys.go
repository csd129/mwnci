package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func FromKeys(args ...object.Object) object.Object {
	if err := typing.Check(
		"fromkeys", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	Value := int64(0)
	Array := args[0].(*object.Array)
	newHash := make(map[object.HashKey]object.HashPair)

	if len(args) == 2 {
		Value = args[1].(*object.Integer).Value
	}

	for _, data := range Array.Elements {
		key := &object.String{Value: fmt.Sprintf("%v", data)}
		val := &object.Integer{Value: int64(Value)}
		newHashPair := object.HashPair{Key: key, Value: val}
		newHash[key.HashKey()] = newHashPair
	}
	return &object.Hash{Pairs: newHash}

}
