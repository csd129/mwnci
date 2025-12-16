package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

// Delete a given hash-key
func harrayDelete(args ...object.Object) object.Object {
     if err := typing.Check(
		"delete", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError("%s", err.Error())
	} 

	if args[0].Type() == object.HASH_OBJ {
		hash := args[0].(*object.Hash)
		key, ok := args[1].(object.Hashable)
		if !ok {
			return newError("key delete() into HASH must be Hashable, got=%s",args[1].Type())
		}
		newHash := make(map[object.HashKey]object.HashPair)
		for k, v := range hash.Pairs {
			if k != key.HashKey() {
				newHash[k] = v
			}
		}
		return &object.Hash{Pairs: newHash}
	}

	if args[0].Type() == object.ARRAY_OBJ {
		arr := args[0].(*object.Array)
		nums := arr.Copy()
		elem := 0
		if args[1].Type() == object.INTEGER_OBJ {
			elem = int(args[1].(*object.Integer).Value)
		} else {
			return newError("argument to delete() must be an INTEGER, got = %s", args[1].Type())
		}
		if (elem > len(arr.Elements)-1) || (elem < 0) {
			return newError("IndexError: array index [%d] out of range ", elem)
		}
		copy(nums.Elements[elem:], nums.Elements[elem+1:])
		len := len(nums.Elements)
		nums.Elements = nums.Elements[:len-1]
		return nums
	}
	return newError("argument to delete() must be HASH or ARRAY, got=%s", args[0].Type())
}

