package evaluator

import (
        "bytes"
        "fmt"
        "encoding/json"
        "sigs.k8s.io/yaml"
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

func jtoy(args ...object.Object) object.Object {
	if err := typing.Check(
		"jsontoyaml", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}

	Stringy := fmt.Sprintf("%v", &object.String{Value: args[0].String()})
	Json := []byte(Stringy)
	Yaml, err := yaml.JSONToYAML(Json)
	if err != nil {
		return newError("%s", err.Error())
	}
	return &object.String{Value: string(Yaml)}
}

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

func ytoj(args ...object.Object) object.Object {
	if err := typing.Check(
		"yamltojson", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}

	Stringy := fmt.Sprintf("%v", &object.String{Value: args[0].String()})
	Yaml := []byte(Stringy)
	Json, err := yaml.YAMLToJSON(Yaml)
	if err != nil {
		return newError("%s", err.Error())
	}
	return &object.String{Value: string(Json)}
}
