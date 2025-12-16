package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"regexp"
)

// regular expression match
func matchFun(args ...object.Object) object.Object {
	if err := typing.Check(
		"match", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	//
	// Compile and match
	//
	reg, err := regexp.Compile(args[0].(*object.String).Value)
	if err != nil {
		return newError("failed to compile regexp %s: %s", args[0].String(), err)
	}
	res := reg.FindStringSubmatch(args[1].(*object.String).Value)
	if len(res) > 0 {
		return TRUE
	}
	return FALSE
}
