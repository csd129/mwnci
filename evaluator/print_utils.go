package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func PrintFun(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Print(arg)
	}
	return NULL
}

func PrintatFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"printat", args,
		typing.MinimumArgs(2),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	Line := args[0].(*object.Integer).Value
	Column := args[1].(*object.Integer).Value
	args = args[2:]
	fmt.Printf("\033[%d;%dH", Line, Column)
	for _, arg := range args {
		fmt.Print(arg)
	}
	return NULL
}

func PrintfFun(args ...object.Object) object.Object {

	// Convert to the formatted version, via our `sprintf`
	// function.
	out := sprintfFun(args...)

	// If that returned a string then we can print it
	if out.Type() == object.STRING_OBJ {
		fmt.Print(out.(*object.String).Value)
	}
	return NULL
}

func PrintlnFun(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Print(arg)
	}
	fmt.Print("\n")
	return NULL
}
