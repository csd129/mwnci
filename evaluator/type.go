package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// type of an item
func typeFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"type", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError("%s", err.Error())
	}
	switch args[0].(type) {
	case *object.String:
		return &object.String{Value: "string"}
	case *object.Regexp:
		return &object.String{Value: "regexp"}
	case *object.Boolean:
		return &object.String{Value: "bool"}
	case *object.Builtin:
		return &object.String{Value: "builtin"}
	case *object.File:
		return &object.String{Value: "file"}
	case *object.Array:
		return &object.String{Value: "array"}
	case *object.Function:
		return &object.String{Value: "function"}
	case *object.Integer:
		return &object.String{Value: "integer"}
	case *object.Float:
		return &object.String{Value: "float"}
	case *object.Hash:
		return &object.String{Value: "hash"}
	default:
		return newError("argument to `type` not supported, got=%s",
			args[0].Type())
	}
}

