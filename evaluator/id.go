package evaluator

import (
	"fmt"

	"mwnci/object"
	"mwnci/typing"
)

// IdOf ...
func IdOf(args ...object.Object) object.Object {
	if err := typing.Check(
		"id", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}

	arg := args[0]

	if n, ok := arg.(*object.Null); ok {
		return &object.String{Value: fmt.Sprintf("%p", n)}
	} else if b, ok := arg.(*object.Boolean); ok {
		return &object.String{Value: fmt.Sprintf("%p", b)}
	} else if i, ok := arg.(*object.Integer); ok {
		return &object.String{Value: fmt.Sprintf("%p", i)}
	} else if s, ok := arg.(*object.String); ok {
		return &object.String{Value: fmt.Sprintf("%p", s)}
	} else if a, ok := arg.(*object.Array); ok {
		return &object.String{Value: fmt.Sprintf("%p", a)}
	} else if h, ok := arg.(*object.Hash); ok {
		return &object.String{Value: fmt.Sprintf("%p", h)}
	} else if f, ok := arg.(*object.Function); ok {
		return &object.String{Value: fmt.Sprintf("%p", f)}
	} else if b, ok := arg.(*object.Builtin); ok {
		return &object.String{Value: fmt.Sprintf("%p", b)}
	}

	return NULL
}
