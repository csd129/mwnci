package evaluator

import (
	"regexp"
	"fmt"

	"mwnci/object"
	"mwnci/typing"
)

// regular expression match
func matchFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"match", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}


	//
	// Compile and match
	//
	reg := regexp.MustCompile(args[0].(*object.String).Value)
	res := reg.FindStringSubmatch(args[1].(*object.String).Value)

	if len(res) > 0 {
		return &object.String{Value: fmt.Sprintf("%v",res)}
	}

	// No match
	return FALSE
}

func init() {
	RegisterBuiltin("match",
		func(env *object.Environment, args ...object.Object) object.Object {return (matchFun(args...))
		})
}
