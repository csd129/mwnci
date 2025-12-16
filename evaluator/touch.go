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
		return newError("%s", err.Error())
	}

	path := args[0].String()

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling touch(): Can't create file %s\n", path)
			return FALSE
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(path, currentTime, currentTime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling touch(): %s - Unable to modify time\n", path)
			return FALSE
		}
	}
	return TRUE
}
