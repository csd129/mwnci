package evaluator

import (
	"fmt"
	"mwnci/object"
	"mwnci/typing"
	"os"
	"time"
)

func Touch(args ...object.Object) object.Object {
	if err := typing.Check(
		"touch", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	path := args[0].Inspect()

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling touch(): %v\n", err)
			return FALSE
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(path, currentTime, currentTime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling touch(): %v\n", err)
			return FALSE
		}
	}
	return TRUE
}
