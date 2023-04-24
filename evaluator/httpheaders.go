package evaluator

import (
	"fmt"
	"net/http"
	"mwnci/object"
	"mwnci/typing"
)

func HttpHeaders(args ...object.Object) object.Object {
	if err := typing.Check(
		"httpheaders", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	url := args[0].(*object.String)
	geturl := fmt.Sprintf("%v", url)
	res, err := http.Get(geturl)
	if err != nil {
		return newError(err.Error())
	}
	defer res.Body.Close()
	var (
		line string
	)
	for k, v := range res.Header {
		line = fmt.Sprintf("%s%s: %s\n", line, k, v)
	}
	return &object.String{Value: line}
}

func init() {
	RegisterBuiltin("httpheaders",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (HttpHeaders(args...))
		})
}
