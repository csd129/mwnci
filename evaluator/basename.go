package evaluator

import (
	"path"
	"mwnci/object"
	"mwnci/typing"
)

// Basename ...
func basenameFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"basename", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}     
	return &object.String{Value: path.Base(args[0].(*object.String).Value)}
}
