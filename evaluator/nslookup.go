package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"net"
)

func NSLookup(args ...object.Object) object.Object {
	if err := typing.Check(
		"nslookup", args,
		typing.RangeOfArgs(0, 2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	domain := args[0].(*object.String).Value
	search_type := "ip"
	if len(args) == 2 {
		search_type = args[1].(*object.String).Value
	}

	switch search_type {
	case "cname":
		record, err := net.LookupCNAME(domain)
		if err != nil {
			return NULL
		}
		return &object.String{Value: record}
	case "host":
		record, err := net.LookupHost(domain)
		if err != nil {
			return NULL
		}
		elements := make([]object.Object, len(record))
		for i, ip := range record {
			elements[i] = &object.String{Value: fmt.Sprint(ip)}
		}
		return &object.Array{Elements: elements}
	case "ip":
		record, err := net.LookupIP(domain)
		if err != nil {
			return NULL
		}
		elements := make([]object.Object, len(record))
		for i, data := range record {
			elements[i] = &object.String{Value: fmt.Sprint(data)}
		}
		return &object.Array{Elements: elements}
	case "txt":
		record, err := net.LookupTXT(domain)
		if err != nil {
			return NULL
		}
		elements := make([]object.Object, len(record))
		for i, data := range record {
			elements[i] = &object.String{Value: fmt.Sprint(data)}
		}
		return &object.Array{Elements: elements}
	case "ptr":
		record, err := net.LookupAddr(domain)
		if err != nil {
			return NULL
		}
		elements := make([]object.Object, len(record))
		for i, data := range record {
			elements[i] = &object.String{Value: fmt.Sprint(data)}
		}
		return &object.Array{Elements: elements}
	case "ns":
		record, _ := net.LookupNS(domain)
		elements := make([]object.Object, len(record))
		for i, data := range record {
			elements[i] = &object.String{Value: fmt.Sprintf("%v", data.Host)}
		}
		return &object.Array{Elements: elements}
	case "mx":
		record, _ := net.LookupMX(domain)
		elements := make([]object.Object, len(record))
		for i, data := range record {
			hostpref := fmt.Sprintf("%v %v", data.Host, data.Pref)
			elements[i] = &object.String{Value: hostpref}
		}
		return &object.Array{Elements: elements}
	default:
		return newError("Search type argument to `nslookup` is not supprted, got `%s`. \nTypes are `cname`, `host`, `ip`, `txt`, `ptr`, `ns`, `mx`. ", search_type)
	}
}
