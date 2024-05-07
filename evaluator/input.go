package evaluator

import (

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
	prompt := "> "
	if len(args) == 1 {
		prompt = args[0].(*object.String).Value
	}
	line := ""
	rl, _ := readline.New(prompt)
	defer rl.Close()
	line, _ = rl.Readline()
	return &object.String{Value: string(line)}
}
