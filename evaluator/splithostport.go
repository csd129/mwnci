package evaluator

import (
	"net"
	"mwnci/object"
	"mwnci/typing"
)

// Change a mode of a file - note the second argument is a string
// to emphasise octal.
func SplitHP(args ...object.Object) object.Object {
	if err := typing.Check(
		"splithostport", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}     


	HostPort := args[0].Inspect()
	host, port, err := net.SplitHostPort(HostPort)
	if err != nil {
		return newError(err.Error())
	}
	elements := make([]object.Object, 2)
	elements[0] = &object.String{Value: host}
	elements[1] = &object.String{Value: port}
	return &object.Array{Elements: elements}

}

