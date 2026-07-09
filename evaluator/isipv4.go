package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"net"
)

func Isipv4(args ...object.Object) object.Object {
	if err := typing.Check(
		"isipv4", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	clientip := args[0].(*object.String).Value

	ip := net.ParseIP(clientip)
	if ip == nil {
		return FALSE
	}
	if ip.To4() != nil {
		return TRUE
	}
	return FALSE
}
