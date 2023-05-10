package evaluator

import (
	"fmt"
	"os"

	"mwnci/object"
	"mwnci/typing"
	"github.com/chzyer/readline"
	
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
	prompt:=""
	line:=""
	if len(args) == 1 {
		prompt := args[0].(*object.String).Value
		fmt.Fprintf(os.Stdout, prompt)
	}

	rl, err := readline.New(string(prompt))
	if err != nil {
		return newError(err.Error())
	}
	defer rl.Close()
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
	return &object.String{Value: string(line)}
	}
	return &object.String{Value: string(line)}
}

func init() {
	RegisterBuiltin("input",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Input(args...))
		})
}
