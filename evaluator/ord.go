package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Ord ...
func Ord(args ...object.Object) object.Object {
	if err := typing.Check(
		"ord", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	s := args[0].(*object.String)
	if len(s.Value) == 1 {
		return &object.Integer{Value: int64(s.Value[0])}
	}
	return newError(
		"TypeError: ord() expected a single character `str` got=%s",
		s.Inspect(),
	)
}
