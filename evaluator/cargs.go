package evaluator

import (
	"regexp"
	"mwnci/object"
	"mwnci/typing"
)

func Cargs(args ...object.Object) object.Object {
	if err := typing.Check(
		"cargs", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	s := args[0].(*object.String).Value
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
	res := r.FindAllString(s, -1)
	elements := make([]object.Object, len(res))
	for i, token := range res {
		elements[i] = &object.String{Value: token}
	}
	return &object.Array{Elements: elements}
}

