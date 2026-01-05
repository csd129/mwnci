package evaluator

import (
	"strconv"
	"fmt"
	"maps"
	"mwnci/object"
	"mwnci/typing"
	"github.com/guptarohit/asciigraph"
)

func Graph(args ...object.Object) object.Object {
	if err := typing.Check(
		"graph", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.ARRAY_OBJ, object.HASH_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	var foo string
	Caption := ""
	Height := int(0)
	Width := int(0)
	newHash := make(map[object.HashKey]object.HashPair)
	Data := args[0].(*object.Array)
        data := make([]float64, len(Data.Elements))
	if len(args) > 1 {
		hash := args[1].(*object.Hash)
		maps.Copy(newHash, hash.Pairs)
	}
		
	for _, ent := range newHash {
		keyname := fmt.Sprintf("%v", ent.Key)
		if keyname == "caption" {
			Caption = fmt.Sprintf("%v", ent.Value)
		}
		if keyname == "height" {
			foo = fmt.Sprintf("%v",ent.Value)
			Height,_ = strconv.Atoi(foo)
		}
		if keyname == "width" {
			foo = fmt.Sprintf("%v",ent.Value)
			Width,_ = strconv.Atoi(foo)
		}
	}
	for i, v := range Data.Elements {
		data[i]=v.(*object.Float).Value
 	} 
       graph := asciigraph.Plot(data, asciigraph.Caption(Caption),
	       asciigraph.Height(Height), asciigraph.Width(Width))

	return &object.String{Value: string(graph)}
}
