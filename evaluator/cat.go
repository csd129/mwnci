package evaluator

import (
	"fmt"
	"os"
	"strings"
	"mwnci/object"
	"mwnci/typing"
)

// Cat ...
func Cat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	data, err := os.ReadFile(filename)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	ndata:=fmt.Sprintf("%s", data)
	ndata=strings.TrimSuffix(ndata, "\n")
	return &object.String{Value: string(ndata)}}
func init() {
	RegisterBuiltin("cat",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Cat(args...))
		})
}
