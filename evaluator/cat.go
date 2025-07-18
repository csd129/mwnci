package evaluator

import (
	"compress/bzip2"
	"compress/gzip"
	"net/http"

	//	"archive/zip"
	"io/ioutil"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"strings"
	//	"fmt"
)

func ZConCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	data, err := os.Open(filename)
	if err != nil {
		return newError(err.Error())
	}
	defer data.Close()
	buff := make([]byte, 512)
	if _, err = data.Read(buff); err != nil {
		return newError(err.Error())
	}
	fileType := http.DetectContentType(buff)
	switch fileType {
	case "application/x-gzip":
		return GzCat(args[0])
		//	case "application/zip":
		//		Content="ZIP"
	default:
		if strings.HasPrefix(string(buff), "\x42\x5a\x68") {
			return BzCat(args[0])
		} else {
			return ConCat(args[0])
		}
	}
}

func ConCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"cat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	data, err := os.ReadFile(args[0].(*object.String).Value)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	return &object.String{Value: string(data)}
}

func GzCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"gzcat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	gzfile, err := os.Open(filename)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	defer gzfile.Close()
	gzread, err := gzip.NewReader(gzfile)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	defer gzread.Close()
	content, err := ioutil.ReadAll(gzread)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	return &object.String{Value: string(content)}
}

func BzCat(args ...object.Object) object.Object {
	if err := typing.Check(
		"bzcat", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	filename := args[0].(*object.String).Value
	bzfile, err := os.Open(filename)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	defer bzfile.Close()
	bzread := bzip2.NewReader(bzfile)
	content, err := ioutil.ReadAll(bzread)
	if err != nil {
		return newError("IOError: error reading from file %s: %s", filename, err)
	}
	return &object.String{Value: string(content)}
}
