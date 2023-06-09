package evaluator

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"mwnci/object"
	"mwnci/typing"
)

func HttpGet(args ...object.Object) object.Object {
	if err := typing.Check(
		"httpget", args,
		typing.RangeOfArgs(1,2),
		typing.WithTypes(object.STRING_OBJ, object.HASH_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	url := args[0].(*object.String)
	geturl := fmt.Sprintf("%v", url)
	client := http.Client{}
	req, err := http.NewRequest("GET", geturl, nil)
	if (err != nil) {
		return newError(err.Error())
	}
	if len(args) == 2 {
		hash:= args[1].(*object.Hash)
		for _, ent := range hash.Pairs {
			hkey := fmt.Sprintf("%v", ent.Key)
			hval := fmt.Sprintf("%v", ent.Value)
			req.Header.Add(hkey, hval)
		}
	}

	res, err := client.Do(req)
	if (err != nil) {
		return newError(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return &object.String{Value: fmt.Sprintf("%s", body)}
}

func init() {
	RegisterBuiltin("httpget",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (HttpGet(args...))
		})
}
