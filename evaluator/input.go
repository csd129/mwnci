package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"github.com/peterh/liner"
)

func Input(args ...object.Object) object.Object {
	if err := typing.Check(
		"input", args,
		typing.RangeOfArgs(0, 1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	prompt := "> "
	if len(args) == 1 {
		prompt = args[0].(*object.String).Value
	}
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	entered, err := line.Prompt(prompt)
	if err == liner.ErrPromptAborted {
		fmt.Printf("Aborted\n")
	}

	return &object.String{Value: string(entered)}
}
