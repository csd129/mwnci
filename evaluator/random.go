package evaluator

import (
	"math/rand"
	"mwnci/object"
	"mwnci/typing"
	"time"
)

// Random ...
func Random(args ...object.Object) object.Object {
	if err := typing.Check(
		"random", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	value := int(args[0].(*object.Integer).Value)
	if value <= 1 {
		return newError("Random range %d is not greater than 1", value)
	}
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return &object.Integer{Value: int64(rng.Intn(value))}
}
