package evaluator

import (
	"math"
	"mwnci/object"
	"mwnci/typing"
)

func Modf(args ...object.Object) object.Object {
	if err := typing.Check(
		"modf", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.FLOAT_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	a := args[0].(*object.Float).Value
	Int, Frac := math.Modf(a)
	elements := make([]object.Object, 2)
	elements[0] = &object.Float{Value: Int}
	elements[1] = &object.Float{Value: Frac}
	return &object.Array{Elements: elements}
}

func init() {
	RegisterBuiltin("modf",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Modf(args...))
		})
}
