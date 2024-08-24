package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Oct ...
func I2IP(args ...object.Object) object.Object {
	if err := typing.Check(
		"inttoip", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	i := args[0].(*object.Integer)
	return &object.String{Value: i.ToIP()}
}






