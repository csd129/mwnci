package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"mwnci/object"
	"mwnci/typing"
	"os"
)

// Input ...
func Input(args ...object.Object) object.Object {
	if err := typing.Check(
		"input", args,
		typing.RangeOfArgs(0, 1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	prompt := "> "
	if len(args) == 1 {
		prompt = args[0].(*object.String).Value
	}
	fmt.Fprintf(os.Stdout, "%s", prompt)
	buffer := bufio.NewReader(os.Stdin)

	line, _, err := buffer.ReadLine()
	if err != nil && err != io.EOF {
		return newError(fmt.Sprintf("error reading input from stdin: %s", err))
	}
	return &object.String{Value: string(line)}
}
