package evaluator

import (
	"fmt"
	"mwnci/lexer"
	"mwnci/object"
	"mwnci/parser"
	"mwnci/typing"
	"strings"
)

func Ucount(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check(
		"ucount", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	var b strings.Builder
	arr := args[0].(*object.Array)
	dict := make(map[string]int)
	b.WriteString("{")
	for _, data := range arr.Elements {
		HashData := fmt.Sprintf("%v", data)
		dict[HashData]++
	}
	counter := 1
	buildline := ""
	dictlen := len(dict)
	for k, v := range dict {
		if counter != dictlen {
			buildline = fmt.Sprintf("\"%v\": %v,", k, v)
		} else {
			buildline = fmt.Sprintf("\"%v\": %v}", k, v)
		}

		fmt.Fprintf(&b, "%v", buildline)
		counter++
	}

	longstring := fmt.Sprintf("%v", b.String())
	l := lexer.New(longstring)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) == 0 {
		return (Eval(program, env))
	}

	return &object.Null{}
}
