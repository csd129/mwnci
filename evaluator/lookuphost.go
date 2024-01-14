package evaluator

import (
	"fmt"
	"net"
	"mwnci/object"
	"mwnci/typing"
)

// LookupHOST ...
func LookupHOST(args ...object.Object) object.Object {
	if err := typing.Check(
		"lookuphost", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	domain := args[0].(*object.String).Value
	record, err := net.LookupHost(domain)
	if err != nil {
		return FALSE
	}
	elements := make([]object.Object, len(record))
	for i, ip := range record {
		elements[i] = &object.String{Value: fmt.Sprint(ip)}
	}
	return &object.Array{Elements: elements}
}
