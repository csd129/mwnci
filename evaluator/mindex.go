package evaluator

import (
	"regexp"
	"index/suffixarray"
	"mwnci/object"
	"mwnci/typing"
)

func Mindex(args ...object.Object) object.Object {
	if err := typing.Check(
		"mindex", args,
		typing.ExactArgs(2),
	); err != nil {
		return newError("%s", err.Error())
	}
	CountArray := make([]object.Object, 0)

	if haystack, ok := args[0].(*object.Array); ok {
		needle := args[1].(object.Comparable)
		for i := range haystack.Elements {
			if needle.Compare(haystack.Elements[i]) == 0 {
				CountArray=append(CountArray, &object.Integer{Value: int64(i)})
			}
		}
	}

	if haystack, ok := args[0].(*object.String); ok {
		if err := typing.Check(
			"index", args,
			typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
		); err != nil {
			return newError("%s", err.Error())
		}
		Results := make([]object.Object, 0)
		needle := args[1].(*object.String).Value
		reg := regexp.MustCompile(needle)
		index := suffixarray.New([]byte(haystack.Value))
		results := index.FindAllIndex(reg, -1)
		for _, entry := range results {
			Results = append(Results, &object.Integer{Value: int64(entry[0])})
		}
		return &object.Array{Elements: Results}
	}
	
	return &object.Array{Elements: CountArray}
}
