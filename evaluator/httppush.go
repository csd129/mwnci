package evaluator

import (
	"bytes"
	"io"
	"mwnci/object"
	"mwnci/typing"
	"net/http"
)

func HttpPost(args ...object.Object) object.Object {
	if err := typing.Check(
		"httppost", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	url := args[0].(*object.String).Value
	content := args[1].(*object.String).Value
	data := args[2].(*object.String).Value

	resp, err := http.Post(url, content, bytes.NewBuffer([]byte(data)))

	if err != nil {
		return newError(err.Error())
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return newError(err.Error())
	}

	return &object.String{Value: string(b)}
}
