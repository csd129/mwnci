package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"net"
	"strings"
)

func Listcidr(args ...object.Object) object.Object {
	if err := typing.Check(
		"listcidrips", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	network := args[0].(*object.String).Value

	ip, ipnet, err := net.ParseCIDR(network)
	if err != nil {
		return newError(err.Error())
	}
	var iplist strings.Builder
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); ipinc(ip) {
		fmt.Fprintf(&iplist, "%v ", ip)
	}
	full_list := fmt.Sprintf("%v", iplist.String())
	trimmed := strings.Trim(full_list, " \n\t\r")
	tokens := strings.Fields(trimmed)
	elements := make([]object.Object, len(tokens))
	for i, token := range tokens {
		elements[i] = &object.String{Value: token}
	}
	return &object.Array{Elements: elements}
}

func ipinc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func init() {
	RegisterBuiltin("listcidrips",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Listcidr(args...))
		})
}
