package evaluator

import (
	"bufio"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"regexp"
)

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

	file, _ := os.Open(filename)
	defer file.Close()
	regarray := make([]object.Object, 0)
	r := regexp.MustCompile(search)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			regarray = append(regarray, &object.String{Value: scanner.Text()})
		}
	}
	return &object.Array{Elements: regarray}
}
