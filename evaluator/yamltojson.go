package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"

	"sigs.k8s.io/yaml"
)

func ytoj(args ...object.Object) object.Object {
	if err := typing.Check(
		"yamltojson", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	Stringy := fmt.Sprintf("%v", &object.String{Value: args[0].String()})
	Yaml := []byte(Stringy)
	Json, err := yaml.YAMLToJSON(Yaml)
	if err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: string(Json)}
}
