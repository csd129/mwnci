package evaluator

import (
	"os"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"mwnci/object"
	"mwnci/typing"
)

func HttpPost(args ...object.Object) object.Object {
	if err := typing.Check(
		"httppost", args,
		typing.ExactArgs(3),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	PROXY:=os.Getenv("HTTP_PROXY")
	proxyURL, _ := url.Parse(PROXY)
	url := args[0].(*object.String).Value
	content := args[1].(*object.String).Value
	data := args[2].(*object.String).Value
	proxy := http.ProxyURL(proxyURL)
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Add("Content-Type", content)
	resp, err := client.Do(req)
	if err != nil {
		return newError(err.Error())
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return newError(err.Error())
	}

	return &object.String{Value: string(b)}
}
