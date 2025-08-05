package evaluator

import (
	//	"fmt"
	"mwnci/object"
	"mwnci/typing"
)

func ArrayInsert(args ...object.Object) object.Object {
	if err := typing.Check(
		"insert", args,
		typing.ExactArgs(3),
		//typing.WithTypes(object.ARRAY_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	if args[0].Type() == object.ARRAY_OBJ {
		arr := args[0].(*object.Array)
		newArray := arr.Copy()
		elem := int(args[1].(*object.Integer).Value)

		if (elem > len(newArray.Elements)-1) || (elem < 0) {
			return newError("IndexError: array index [%d] out of range ", elem)
		} else {
			val := args[2]
			newArray.Insert(elem, val)
		}
		return newArray
	}
	if args[0].Type() == object.STRING_OBJ {
		text := args[0].(*object.String).Value
		elem := int(args[1].(*object.Integer).Value)
		val:=args[2].(*object.String).Value
		if (elem > len(text)-1) || (elem < 0) {
			return newError("IndexError: string index [%d] out of range ", elem)
		}
		Newtext := text[:elem] + val  + text[elem:]
		return &object.String{Value: Newtext} 
	}
	return newError("argument to insert() not supported, expected HASH, ARRAY or STRING, got=%s", args[0].Type())

}

