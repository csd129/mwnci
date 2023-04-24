package evaluator

import (
	"fmt"

	"mwnci/object"
	"mwnci/typing"
)

// Chr ...
func Chr(args ...object.Object) object.Object {
	if err := typing.Check(
		"chr", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}     

	i := args[0].(*object.Integer)
	return &object.String{Value: fmt.Sprintf("%c", rune(i.Value))}
}

func init() {
	RegisterBuiltin("chr",
		func(env *object.Environment, args ...object.Object) object.Object {
			return(Chr(args...))
		})
}
