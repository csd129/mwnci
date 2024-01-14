package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"net"
)

func Listcidrips(args ...object.Object) object.Object {
	if err := typing.Check(
		"listcidrips", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	network := args[0].(*object.String).Value
	ip, ipnet, err := net.ParseCIDR(network)
	ip_array := make([]object.Object, 0)
	if err != nil {
		return newError(err.Error())
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); IncIP(ip) {
		ip_array = append(ip_array, &object.String{Value: fmt.Sprint(ip)})
	}
	return &object.Array{Elements: ip_array}
}

func IncIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

