package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func SecToTime(args ...object.Object) object.Object {
	if err := typing.Check(
		"sec2time", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	Seconds := args[0].(*object.Integer).Value
	Day := &object.Integer{Value: int64(Seconds/60/60/24)}
	Hour := &object.Integer{Value: int64((Seconds/60/60)%24)}
	Min := &object.Integer{Value: int64((Seconds/60)%60)}
	Sec := &object.Integer{Value: int64(Seconds%60)}

	newHash := make(map[object.HashKey]object.HashPair)
	key:=&object.String{Value: "Day"}
	newHashPair := object.HashPair{Key: key, Value: Day}
	newHash[key.HashKey()] = newHashPair
	key=&object.String{Value: "Hour"}
	newHashPair = object.HashPair{Key: key, Value: Hour}
	newHash[key.HashKey()] = newHashPair
	key=&object.String{Value: "Min"}
	newHashPair = object.HashPair{Key: key, Value: Min}
	newHash[key.HashKey()] = newHashPair
	key=&object.String{Value: "Sec"}
	newHashPair = object.HashPair{Key: key, Value: Sec}
	newHash[key.HashKey()] = newHashPair

	return &object.Hash{Pairs: newHash}
}
