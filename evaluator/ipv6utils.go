package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"net/netip"
)

func Expandv6(args ...object.Object) object.Object {
	if err := typing.Check(
		"expandv6", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	IPv6Address := args[0].(*object.String).Value
	Expanded, _ := netip.ParseAddr(IPv6Address)
	FullIP := fmt.Sprintf("%v", Expanded.StringExpanded())
	return &object.String{Value: string(FullIP)}
}

func Shrinkv6(args ...object.Object) object.Object {
	if err := typing.Check(
		"shrinkv6", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	IPv6Address := args[0].(*object.String).Value
	Shrink, _ := netip.ParseAddr(IPv6Address)
	Shrunk := fmt.Sprintf("%v", Shrink.String())
	return &object.String{Value: string(Shrunk)}
}
