package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strconv"
	"strings"
)

func IP2I(args ...object.Object) object.Object {
	IPInt := 0
	if err := typing.Check(
		"iptoint", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}
	switch args[0].(type) {
	case *object.String:
		ip := args[0].(*object.String).Value
		if CheckIP(args[0].(*object.String)) == FALSE {
			return args[0]
		}
		tokens := strings.Split(ip, ".")
		octets := make([]int, len(tokens))
		for i, token := range tokens {
			octets[i], _ = strconv.Atoi(token)
		}
		IPInt = octets[3] + (octets[2] << 8) + (octets[1] << 16) + (octets[0] << 24)
		return &object.Integer{Value: int64(IPInt)}
	case *object.Array:
		arr := args[0].(*object.Array)
		newArray := arr.Copy()
		for i, v := range newArray.Elements {
			newArray.Aset(i, IP2I(v))
		}
		return newArray
	default:
		return args[0]
	}
}
