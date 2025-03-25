package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
	"fmt"
	"github.com/itchyny/json2yaml"
)

func j2yaml(args ...object.Object) object.Object {
	if err := typing.Check(
		"toyaml", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}

	//	Json := args[0].(*object.String).Value
	Json := fmt.Sprintf("%v", &object.String{Value: args[0].String()})
	input := strings.NewReader(Json)
	var output strings.Builder
	if err := json2yaml.Convert(&output, input); err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: output.String()}
}
