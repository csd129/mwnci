package evaluator

import (
	"net"
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func CheckIP(args ...object.Object) object.Object {
	if err := typing.Check(
		"checkip", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	clientip := args[0].(*object.String).Value

	ip := fmt.Sprintf("%v", net.ParseIP(clientip))
	
	if ip != clientip {
		return FALSE
	}
	return &object.String{Value: clientip}
}

