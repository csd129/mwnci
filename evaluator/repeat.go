package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)


func Repeat(args ...object.Object) object.Object {
	if err := typing.Check(
		"repeat", args,
		typing.RangeOfArgs(2,3),
		typing.WithTypes(object.STRING_OBJ, object.INTEGER_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	sep := ""
	s := args[0].(*object.String).Value
	l := args[1].(*object.Integer).Value
	long_string := ""
	if len(args) == 3 {
		sep = args[2].(*object.String).Value
	}
	i := int64(1)
	for i <= l {
		if i < l {
			long_string = long_string + s + sep
		} else {
			long_string = long_string + s
		}
		i++
	}
	return &object.String{Value: long_string}
	
}

func init() {
	RegisterBuiltin("repeat",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Repeat(args...))
		})
}
