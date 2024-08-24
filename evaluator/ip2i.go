package evaluator

import (
	"strings"
	"strconv"
	"mwnci/object"
	"mwnci/typing"
)

func IP2I(args ...object.Object) object.Object {
	IPInt := 0
	if err := typing.Check(
		"iptoint", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	ip := args[0].(*object.String).Value
	tokens := strings.Split(ip, ".")
	octets := make([]int, len(tokens))
	for i, token := range tokens {
		octets[i], _ = strconv.Atoi(token)
	}
	IPInt = octets[3] + (octets[2] << 8) + (octets[1] << 16) + (octets[0] << 24)
	return &object.Integer{Value: int64(IPInt)}
}






