package evaluator

import (
	"math/rand"
	"mwnci/object"
	"mwnci/typing"
	"time"
)

func RandSample(args ...object.Object) object.Object {
	if err := typing.Check(
		"randsample", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	RandRange := int(args[0].(*object.Integer).Value)
	ArraySize := int64(args[1].(*object.Integer).Value)
	if RandRange <= 1 {
		return newError("Random range %d is not greater than 1", RandRange)
	}
	if ArraySize < 1 {
		return newError("Array size %d is not greater than 0", ArraySize)
	}

	RandArray := make([]object.Object, ArraySize)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for counter := range RandArray {
		RandArray[counter] = &object.Integer{Value: int64(rng.Intn(RandRange))}
	}
	return &object.Array{Elements: RandArray}
}
