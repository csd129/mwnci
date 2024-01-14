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

	switch args[0].(type) {
	case *object.String:
		txt := args[0].(*object.String).Value

		// Lex the input
		l := lexer.New(txt)

		// parse it.
		p := parser.New(l)

		// If there are no errors
		program := p.ParseProgram()
		if len(p.Errors()) == 0 {
			// evaluate it, and return the output.
			return (EvalContext(context.Background(), program, env))
		}

		// Otherwise abort.  We should have try { } catch
		// to allow this kind of error to be caught in the future!
		fmt.Printf("Error parsing eval-string: %s", txt)
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}
	return newError("argument to `eval` not supported, got %s", args[0].Type())
}

