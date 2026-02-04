package evaluator

import (
	"mwnci/object"
	"mwnci/typing"
	"regexp"
)

// regular expression match & split
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
		return newError("failed to compile regexp %s: %s",
			args[0].Inspect(), err)
	}
	res := reg.FindStringSubmatch(args[1].(*object.String).Value)

	if len(res) > 0 {

		newHash := make(map[object.HashKey]object.HashPair)

		//
		// If we get a match then the output is an array
		// First entry is the match, any additional parts
		// are the capture-groups.
		//
		if len(res) > 1 {
			for i := 1; i < len(res); i++ {

				// Capture groups start at index 0.
				k := &object.Integer{Value: int64(i - 1)}
				v := &object.String{Value: res[i]}

				newHashPair := object.HashPair{Key: k, Value: v}
				newHash[k.HashKey()] = newHashPair

			}
		}

		return &object.Hash{Pairs: newHash}
	}

	// No match
	return NULL
}
