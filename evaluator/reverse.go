package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
)

func Rev(args ...object.Object) object.Object {
	if err := typing.Check(
		"reverse", args,
		typing.ExactArgs(1),
	); err != nil {
		return newError(err.Error())
	}
	
	switch args[0].(type) {
	case *object.String:
		input := args[0].(*object.String).Value
		n := 0
		rune := make([]rune, len(input))
		for _, r := range input { 
			rune[n] = r
			n++
		} 
		rune = rune[0:n]
		// Reverse 
		for i := 0; i < n/2; i++ { 
			rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
		} 
		// Convert back to UTF-8. 
		output := string(rune)
		return &object.String{Value: output}
	case *object.Array:
		foo := args[0].(*object.Array)
		nums := foo.Copy()
		nums.Reverse()
		return nums
	}
	return FALSE
}

