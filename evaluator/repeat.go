package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"strings"
)

func Repeat(args ...object.Object) object.Object {
	if err := typing.Check(
		"repeat", args,
		typing.RangeOfArgs(2, 3),
		typing.WithTypes(object.STRING_OBJ, object.INTEGER_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	sep := ""
	s := args[0].(*object.String).Value
	l := args[1].(*object.Integer).Value
	var b strings.Builder
	if len(args) == 3 {
		sep = args[2].(*object.String).Value
	}
	i := int64(1)
	for i <= l {
		if i < l {
			fmt.Fprintf(&b, "%v%v", s, sep)
		} else {
			fmt.Fprintf(&b, "%v", s)
		}
		i++
	}
	return &object.String{Value: b.String()}

}

