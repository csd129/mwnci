package evaluator

import (
	"context"
	"fmt"
	"os"

	"mwnci/lexer"
	"mwnci/object"
	"mwnci/parser"
	"mwnci/typing"
)

// evaluate a string containing monkey-code
func evalFun(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check(
		"eval", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	txt := args[0].(*object.String).Value
	if len(txt) == 0 {
		return newError("argument to `eval` not supported, got an empty string")
	}

	// Lex the input
	l := lexer.New(txt)

	// parse it.
	p := parser.New(l)

	// If there are no errors
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Printf("Error parsing eval-string: %s", txt)
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}
	return (EvalContext(context.Background(), program, env, false))
}
