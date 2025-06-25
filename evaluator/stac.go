
package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
	"io"
)

func StrRevCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"stac", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	RevString := args[0].(*object.String).Value
	scanner := NewScan(strings.NewReader(RevString), len(RevString))
	var StringBuffer strings.Builder
	for {
		line, _, err := scanner.Line()
		if err == io.EOF {
			break
		}
		StringBuffer.WriteString(string(line) + "\n")
	}
	Output:=strings.TrimSuffix(StringBuffer.String(), "\n")
	return &object.String{Value: Output}
}
