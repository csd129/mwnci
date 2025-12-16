package evaluator

import (
	"io"
	"mwnci/object"
	"mwnci/typing"
	"os"
)

func FCp(args ...object.Object) object.Object {
	if err := typing.Check(
		"cp", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	orig := args[0].String()
	dest := args[1].String()

	source, err := os.Open(orig)
	if err != nil {
		return newError("%s", err.Error())
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return newError("%s", err.Error())
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return newError("%s", err.Error())
	}
	return &object.Integer{Value: nBytes}
}
