package evaluator

import (
	"net"
	"mwnci/object"
	"mwnci/typing"
)

// LookupCNAME ...
func LookupCNAME(args ...object.Object) object.Object {
	if err := typing.Check(
		"lookupcname", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	domain := args[0].(*object.String).Value
	record, err := net.LookupCNAME(domain)
	if err != nil {
		return FALSE
	}
	return &object.String{Value: record}
}

func init() {
	RegisterBuiltin("lookupcname",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (LookupCNAME(args...))
		})
}
