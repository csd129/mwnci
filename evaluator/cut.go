package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"fmt"
	"strings"
)

func FuncCut(args ...object.Object) object.Object {
	if err := typing.Check(
		"cut", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	line := args[0].(*object.String).Value
	pivot := args[1].(*object.String).Value
	elements := make([]object.Object, 2)
	before, after, _ := strings.Cut(line, pivot)

	elements[0] = &object.String{Value: fmt.Sprintf("%v", before)}
	elements[1] = &object.String{Value: fmt.Sprintf("%v", after)}
	return &object.Array{Elements: elements}
}
