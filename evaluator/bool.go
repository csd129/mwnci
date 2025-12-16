package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Bool ...
func Bool(args ...object.Object) object.Object {
	if err := typing.Check(
		"bool", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}     
	retval:=false
	switch args[0].(type) {
	case *object.String:
		value:=args[0].(*object.String).Value
		if len(value) != 0 {
			retval=true
		}
	case *object.Boolean:
		retval=args[0].(*object.Boolean).Value
	case *object.Array:
		foo:=args[0].(*object.Array)
		objarray:=len(foo.Elements)
		if objarray != 0 {
			retval=true
		}
	case *object.Integer:
		value:=args[0].(*object.Integer).Value
		if value != 0 {
			retval=true
		}
	case *object.Float:
		value:=args[0].(*object.Float).Value
		if value != 0 {
			retval=true
		}
	case *object.Hash:
		if size, ok := args[0].(object.Sizeable); ok {
			if size.Len() != 0 {
				retval=true
			}
		}
	default:
		retval=false
	}

	return &object.Boolean{Value: retval}
}
