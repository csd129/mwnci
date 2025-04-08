package evaluator

import (
	"bytes"
	"io"
	"mwnci/object"
	"mwnci/typing"
	"net/http"
	"net/url"
	"os"
	"time"
)

func HttpPostNoProxy(args ...object.Object) object.Object {
	url := args[0].(*object.String).Value
	content := args[1].(*object.String).Value
	data := args[2].(*object.String).Value
	resp, err := http.Post(url, content, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return newError(err.Error())
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return newError(err.Error())
	}
	return &object.String{Value: string(b)}
}

func HttpPost(args ...object.Object) object.Object {
	timeout := 10 // default to 10 second timeout
	if err := typing.Check(
		"httppost", args,
		typing.RangeOfArgs(3, 4),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ, object.STRING_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError(err.Error())
	}
	PROXY := os.Getenv("HTTP_PROXY")
	if len(PROXY) == 0 {
		return HttpPostNoProxy(args[0], args[1], args[2])
	}
	proxyURL, _ := url.Parse(PROXY)
	url := args[0].(*object.String).Value
	content := args[1].(*object.String).Value
	data := args[2].(*object.String).Value
	proxy := http.ProxyURL(proxyURL)
	if len(args) == 4 {
		timeout = int(args[3].(*object.Integer).Value)
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeout) * time.Second,
	}
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
