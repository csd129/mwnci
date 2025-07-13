package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// set a hash-field
func setFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"set", args,
		typing.ExactArgs(3),
	); err != nil {
		return newError(err.Error())
	}
	if args[0].Type() == object.HASH_OBJ {
		key, ok := args[1].(object.Hashable)
		if !ok {
			return newError("key `set` into HASH must be Hashable, got=%s",
				args[1].Type())
		}
		newHash := make(map[object.HashKey]object.HashPair)
		hash := args[0].(*object.Hash)
		for k, v := range hash.Pairs {
			newHash[k] = v
		}
		newHashKey := key.HashKey()
		newHashPair := object.HashPair{Key: args[1], Value: args[2]}
		newHash[newHashKey] = newHashPair
		return &object.Hash{Pairs: newHash}
	}
	if args[0].Type() == object.ARRAY_OBJ {
		arr := args[0].(*object.Array)
//		newArray := arr.Copy()
		elem := int(args[1].(*object.Integer).Value)
		if (elem > len(arr.Elements)-1) || (elem < 0) {
			return newError("IndexError: array index [%d] out of range ", elem)
		} else {
			val := args[2]
			arr.Aset(elem, val)
		}
		return arr
	}
	if args[0].Type() == object.STRING_OBJ {
		text := args[0].(*object.String).Value
		texty := make([]byte, len(text))
		val := args[2].(*object.String).Value
		if len(val) != 1 {
			return newError("StringError: Inserted string must be one character")
		}
		elem := int(args[1].(*object.Integer).Value)
		if (elem > len(text)-1) || (elem < 0) {
			return newError("IndexError: string index [%d] out of range ", elem)
		} else {
			for k, v := range []byte(text) {
				texty[k] = byte(v)
			}
		}
		texty[elem] = val[0]
		return &object.String{Value: string(texty[:])}
	}
	return newError("argument to set() not supported, expected HASH, ARRAY or STRING, got=%s", args[0].Type())
}
