package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Range(args ...object.Object) object.Object {
	if err := typing.Check(
		"range", args,
		typing.RangeOfArgs(1, 3),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	Start := int64(0)
	End := int64(0)
	Step := int64(1)
	Rev := false

	if len(args) == 1 {
		End = args[0].(*object.Integer).Value
	}
	if len(args) > 1 {
		Start = args[0].(*object.Integer).Value
		End = args[1].(*object.Integer).Value
	}
	if len(args) > 2 {
		Step = args[2].(*object.Integer).Value
	}
	if Start > End && Step < 0 {
		Rev = true
	}
	list := make([]object.Object, 0)

	if Rev {
		for i := Start; i > End; i += Step {
			list = append(list, &object.Integer{Value: i})
		}
	} else {
		for i := Start; i < End; i += Step {
			list = append(list, &object.Integer{Value: i})
		}
	}
	return &object.Array{Elements: list}
}
