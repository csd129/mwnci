package evaluator

import (
	"fmt"
	"io"
	"mwnci/object"
	"mwnci/typing"
	"net/http"
)

func HttpGet(args ...object.Object) object.Object {
	if err := typing.Check(
		"httpget", args,
		typing.RangeOfArgs(1, 3),
		typing.WithTypes(object.STRING_OBJ, object.HASH_OBJ, object.HASH_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}
	url := args[0].(*object.String)
	geturl := fmt.Sprintf("%v", url)
	client := http.Client{}
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return newError("%s", err.Error())
	}
	if len(args) == 2 {
		hash := args[1].(*object.Hash)
		for _, ent := range hash.Pairs {
			hkey := fmt.Sprintf("%v", ent.Key)
			hval := fmt.Sprintf("%v", ent.Value)
			req.Header.Add(hkey, hval)
		}
	}

	if len(args) == 3 {
		hash := args[2].(*object.Hash)
		for _, ent := range hash.Pairs {
			hkey := fmt.Sprintf("%v", ent.Key)
			hval := fmt.Sprintf("%v", ent.Value)
			req.AddCookie(&http.Cookie{Name: hkey, Value: hval})
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return newError("%s", err.Error())
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return newError("%s", err.Error())
	}
	return &object.String{Value: string(body)}
}

