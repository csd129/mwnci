
package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"strings"
	"io"
	"os"
)

func RevCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"tac", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	file, err := os.Open(filename)
	if err != nil {
		return newError(err.Error())
	}
	fileStatus, err := file.Stat()
	if err != nil {
		return newError(err.Error())
	}
	defer file.Close()
	scanner := NewScan(file, int(fileStatus.Size()))
	var StringBuffer strings.Builder
	for {
		line, _, err := scanner.LineBytes()
		if err == io.EOF {
			break
		}
		if err != nil  {
			return newError(err.Error())
		}
		StringBuffer.WriteString(string(line) + "\n")
	}
	Output:=strings.TrimSuffix(StringBuffer.String(), "\n")
	return &object.String{Value: Output}
}
