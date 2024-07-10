package evaluator

import (
	"fmt"
	"mwnci/object"
)

func putsFun(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Print(arg)
	}
	return NULL
}
