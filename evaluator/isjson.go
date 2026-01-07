package evaluator

import (
	//	"bytes"
	"encoding/json"
	"mwnci/object"
	"mwnci/typing"
)

func Isjson(args ...object.Object) object.Object {
	if err := typing.Check(
		"isjson", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}

	JsonData := args[0].String()
	var js interface{}
	if err := json.Unmarshal([]byte(JsonData), &js); err != nil {
		return FALSE
	}
	return TRUE
}
