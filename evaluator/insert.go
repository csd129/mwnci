package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func ArrayInsert(args ...object.Object) object.Object {
	if err := typing.Check(
		"insert", args,
		typing.ExactArgs(3),
	); err != nil {
		return newError("%s", err.Error())
	}

	if args[0].Type() == object.ARRAY_OBJ {
		arr := args[0].(*object.Array)
		elem := int(args[1].(*object.Integer).Value)
		if (elem > len(arr.Elements)-1) || (elem < 0) {
			return newError("IndexError: array index [%d] out of range ", elem)
		}
		arr.Insert(elem, args[2])
		return arr
	}
	if args[0].Type() == object.STRING_OBJ {
		text := args[0].(*object.String).Value
		elem := int(args[1].(*object.Integer).Value)
		val := args[2].String()
		if (elem > len(text)-1) || (elem < 0) {
			return newError("IndexError: string index [%d] out of range ", elem)
		}
		Newtext := text[:elem] + val + text[elem:]
		return &object.String{Value: Newtext}
	}
	return newError("argument to insert() not supported, expected ARRAY or STRING, got=%s", args[0].Type())

}
