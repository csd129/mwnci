package evaluator

import (
	"math/rand"
	"time"
	"mwnci/object"
	"mwnci/typing"
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
	i := args[0].(*object.Integer)
	value := int(i.Value)
	if value <= 1 {
		return newError("Value %d is not greater than 1", value)
	}
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(value)
	return &object.Integer{Value: int64(val)}
}
func init() {
	RegisterBuiltin("random",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Random(args...))
		})
}
