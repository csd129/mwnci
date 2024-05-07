package evaluator

import (
	"net"

	"mwnci/object"
	"mwnci/typing"
)

func Ipincidr(args ...object.Object) object.Object {
	if err := typing.Check(
		"ipincidr", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	network := args[0].(*object.String).Value
	clientip := args[1].(*object.String).Value

	_, subnet, _ := net.ParseCIDR(network)
	ip := net.ParseIP(clientip)
	if subnet.Contains(ip) {
		return &object.String{Value: string(clientip)}
	}
	return FALSE
}

