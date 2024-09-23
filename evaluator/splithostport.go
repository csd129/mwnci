package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"net"
)

func SplitHP(args ...object.Object) object.Object {
	if err := typing.Check(
		"splithostport", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	Data := args[0].(*object.String).Inspect()

	host, port, err := net.SplitHostPort(Data)
	if err != nil {
		return newError(err.Error())
	}
	elements := make([]object.Object, 2)
	elements[0] = &object.String{Value: host}
	elements[1] = &object.String{Value: port}
	return &object.Array{Elements: elements}
}
