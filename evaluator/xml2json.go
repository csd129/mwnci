package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"

	xj "github.com/txix-open/goxml2json"
)

func Funxml2json(args ...object.Object) object.Object {
	if err := typing.Check(
		"xmltojson", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}
	xmlData := args[0].(*object.String).Value
	xml := strings.NewReader(xmlData)
	converter := xj.NewConverter()
	json, err := converter.Convert(xml)
	if err != nil {
		return newError("%s", err.Error())
	}
	return &object.String{Value: json.String()}
}

//fmt.Println(json.String())
// {"hello": ["world"]}
