package evaluator

import (
	"os"
	"time"
	
	"mwnci/object"
	"mwnci/typing"
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
			return FALSE
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(path, currentTime, currentTime)
		if err != nil {
			return FALSE
		}
	}
	return &object.Integer{Value: 0}
}

func init() {
	RegisterBuiltin("touch",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (Touch(args...))
		})
}
