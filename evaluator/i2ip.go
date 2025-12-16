package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func I2IP(args ...object.Object) object.Object {
	if err := typing.Check(
		"inttoip", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}
	switch args[0].(type) {
	case *object.Integer:
		i := args[0].(*object.Integer)
		return &object.String{Value: i.ToIP()}
	case *object.Array:
		arr := args[0].(*object.Array)
		newArray := arr.Copy()
		for i, v := range newArray.Elements {
			newArray.Aset(i, I2IP(v))
		}
		return newArray
	default:
		return args[0]
	}
}
