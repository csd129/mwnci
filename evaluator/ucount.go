package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/lexer"
	"mwnci/parser"
	"mwnci/typing"
)

func Ucount(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check(
		"ucount", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.ARRAY_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	//	dict := make(map[string]int)
	arr := args[0].(*object.Array)
	dict := make(map[string]int)
	longstring:="{"
	for _,data := range arr.Elements{
		HashData := fmt.Sprintf("%v", data)
		dict[HashData] = dict[HashData] + 1
	}
	counter := 1
	buildline:=""
	dictlen:=len(dict)
	for k, v := range dict {
		if counter != dictlen {
			buildline = fmt.Sprintf("\"%v\": %v,", k, v)
		} else {
			buildline = fmt.Sprintf("\"%v\": %v}", k, v)
		}
		longstring=fmt.Sprintf("%v%v", longstring, buildline)
		counter++
	}
	l := lexer.New(string(longstring))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) == 0 {
		return (Eval(program, env))
	}
	
	return FALSE
}

func init() {
	RegisterBuiltin("yewcount",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Ucount(env, args...))
		})
}
