package evaluator

import (
	"strings"
	"time"
	"mwnci/object"
	"mwnci/typing"
)

// Time...
func Time(args ...object.Object) object.Object {
	if err := typing.Check(
		"time", args,
		typing.RangeOfArgs(0,1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	now := time.Now()
	if len(args) == 1  {
		TimeType := strings.ToLower(args[0].(*object.String).Value)
		switch TimeType {
		case "milli" :
			return &object.Integer{Value: now.UnixMilli()}
		case "micro" :
			return &object.Integer{Value: now.UnixMicro()}
		case "nano" :
			return &object.Integer{Value: now.UnixNano()}
		case "unix" :
			return &object.Integer{Value: now.Unix()}
		default:
			return newError("argument to time() not supported, got=%s", TimeType)
		}
	}
	return &object.Integer{Value: now.Unix()}
}
