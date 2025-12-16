package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"time"

	"github.com/ncruces/go-strftime"
)

// AscTime...
func AscTime(args ...object.Object) object.Object {
	if err := typing.Check(
		"asctime", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.INTEGER_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	format := ""
	if len(args) == 1 {
		format = "%a %b %d %H:%M:%S %Y"
	} else {
		format = args[1].(*object.String).Value
	}
	timer := time.Unix(args[0].(*object.Integer).Value, 0)
	return &object.String{Value: strftime.Format(format, timer)}

}
