package evaluator

import (
	"encoding/csv"
	"mwnci/object"
	"mwnci/typing"
        "strings"
)

func SplitCsv(args ...object.Object) object.Object {
	if err := typing.Check(
		"splitcsv", args,
		typing.RangeOfArgs(1,2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	Separator := ","
	s := args[0].(*object.String).Value
	if len(args) == 2 {
		Separator = args[1].(*object.String).Value
	}
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = []rune(Separator)[0]
	records, err := r.ReadAll()
	if err != nil {
		return newError("%s", err.Error())
	}
	Main :=  make([]object.Object, len(records))
	Sub :=  make([]object.Object, len(records[0]))

	for i, d := range records {
		for I, D := range d {
			Sub[I] = &object.String{Value: D}
		}
                pushlist := &object.Array{Elements: Sub}
                foo := pushlist.Copy()
		Main[i] = foo
	}
	
	return &object.Array{Elements: Main}
}

