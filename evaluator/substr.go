package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func subStrFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"substr", args,
		typing.RangeOfArgs(2, 3),
		typing.WithTypes(object.STRING_OBJ, object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	text := args[0].(*object.String).Value
	i := int(args[1].(*object.Integer).Value)
	if i < 0 {
		i += len(text)
	}
	l := 0
	if len(args) == 3 {
		l = int(args[2].(*object.Integer).Value)
		if l < 1 {
			l = 1
		}
	} else {
		l = len(text) - i
	}
	asRunes := []rune(text)
	if i == len(asRunes)-1 {
		return &object.String{Value: string(asRunes[i])}
	}
	if i+l > len(asRunes)-1 {
		l = len(asRunes) - i
	}
	return &object.String{Value: string(asRunes[i : i+l])}
}
