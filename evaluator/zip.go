package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func zipFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"zip", args,
		typing.MinimumArgs(2),
	); err != nil {
		return newError("%s", err.Error())
	}

	shortest := len(args[0].(*object.Array).Elements)
	for i := 0; i < len(args); i++ {
		if args[i].Type() == object.ARRAY_OBJ {
			length := len(args[i].(*object.Array).Elements)
			if length < shortest {
				shortest = length
			}
		} else {
			return newError("argument to zip() not supported, expected ARRAY, got=%s", args[i].Type())
		}
	}

	ziparray := make([]object.Object, shortest)
	length:=len(args)
	joinarray := make([]object.Object, length)

	counter := 0

	for counter < shortest {
		for i:=0; i < len(args); i++ {
			joinarray[i] = args[i].(*object.Array).Elements[counter]
		}
		pushlist := &object.Array{Elements: joinarray}
		foo:=pushlist.Copy()
		ziparray[counter]=foo
		counter++
	}
	return &object.Array{Elements: ziparray}
}

