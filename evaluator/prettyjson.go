package evaluator

import (
	"bytes"
	"encoding/json"
	"mwnci/object"
	"mwnci/typing"
)

func Pjson(args ...object.Object) object.Object {
	if err := typing.Check(
		"pjson", args,
		typing.RangeOfArgs(1, 2),
	); err != nil {
		return newError("%s", err.Error())
	}
	Indent := 2
	IndentString := ""
	if len(args) > 1 {
		Indent = int(args[1].(*object.Integer).Value)
	}
	for n := 1; n <= Indent; n++ {
		IndentString = IndentString + " "
	}

	JsonData := args[0].String()

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, []byte(JsonData), "", IndentString)

	return &object.String{Value: prettyJSON.String()}
}
