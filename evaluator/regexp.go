package evaluator

import (
	"bufio"
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"regexp"
)

// ReadFile ...
func ReadFileLines(path string, search string) ([]string, error) {
	file, _ := os.Open(path)
	defer file.Close()
	var lines []string
	r, _ := regexp.Compile(search)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			lines = append(lines, scanner.Text())
		}
	}
	return lines, scanner.Err()
}

func Regexp(args ...object.Object) object.Object {
	if err := typing.Check(
		"regexp", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	filename := args[1].(*object.String).Value
	search := args[0].(*object.String).Value
	lines, _ := ReadFileLines(filename, search)
	elements := make([]object.Object, len(lines))
	for n, line := range lines {
		elements[n] = &object.String{Value: fmt.Sprint(line)}
	}
	return &object.Array{Elements: elements}
}

