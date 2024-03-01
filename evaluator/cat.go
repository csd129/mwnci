package evaluator
import (
	"mwnci/object"
	"mwnci/typing"
	"os"
	"strings"
	"bufio"
	"fmt"
)

// Cat ...
func Cat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	var b strings.Builder
	filename := args[0].(*object.String).Value
	data:=""
	if filename != "-" {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling cat(): %v\n", err.Error())
			return &object.String{Value: string(data)}
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Fprintf(&b, "%v\n", scanner.Text())
		}
		data=fmt.Sprintf("%v", b.String())

	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintf(&b, "%v\n", scanner.Text())
		}
		data=fmt.Sprintf("%v", b.String())
	}
	return &object.String{Value: string(data)}
}
