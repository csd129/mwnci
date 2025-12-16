package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"

	"sigs.k8s.io/yaml"
)

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
