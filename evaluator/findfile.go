package evaluator

import (
	"log"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"path/filepath"
	"regexp"
)

func FindFile(args ...object.Object) object.Object {
	var files []string
	if err := typing.Check(
		"findfile", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	rootpath := args[0].(*object.String).Value
	filesearch := args[1].(*object.String).Value
	r, err := regexp.Compile(filesearch)
	if err != nil {
		return newError("failed to compile regexp %s: %s", filesearch, err)
	}

	err = filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		basefile := filepath.Base(path)
		if r.MatchString(basefile) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	elements := make([]object.Object, len(files))
	if len(files) != 0 {
		for i, token := range files {
			elements[i] = &object.String{Value: token}
		}
	}
	return &object.Array{Elements: elements}
}
