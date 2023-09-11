package evaluator

import (
	"time"
	"mwnci/object"
	"mwnci/typing"
	"github.com/ncruces/go-strftime"
)

// AscTime...
func AscTime(args ...object.Object) object.Object {
	if err := typing.Check(
		"asctime", args,
		typing.RangeOfArgs(1,2),
		typing.WithTypes(object.INTEGER_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	format:=""
	if len(args) == 1 {
		format="%a %b %d %H:%M:%S %Y"
	} else {
		format=args[1].(*object.String).Value
	}
	epochtime := args[0].(*object.Integer).Value
	timer := time.Unix(epochtime, 0)
	formatted := strftime.Format(format, timer)
	return &object.String{Value: formatted}

}

func init() {
	RegisterBuiltin("asctime",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (AscTime(args...))
		})
}
