package evaluator

import (
	"encoding/json"
	"encoding/xml"
	"mwnci/object"
	"mwnci/typing"
	"strings"

	"vimagination.zapto.org/json2xml"
)

func Funjson2xml(args ...object.Object) object.Object {
	if err := typing.Check(
		"jsontoxml", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}
	jsonData := args[0].(*object.String).Value
	var buf strings.Builder
	x := xml.NewEncoder(&buf)
	x.Indent("", "\t")
	if err := json2xml.Convert(json.NewDecoder(strings.NewReader(jsonData)), x); err != nil {
		return newError("%s", err.Error())
	}
	x.Flush()
	return &object.String{Value: buf.String()}
}
