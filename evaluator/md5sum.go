package evaluator

import (
	"crypto/md5"
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func Md5sum(args ...object.Object) object.Object {
	if err := typing.Check(
		"md5sum", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	data := []byte(args[0].(*object.String).Value)
	sum := fmt.Sprintf("%x", md5.Sum(data))
	return &object.String{Value: sum}
}
