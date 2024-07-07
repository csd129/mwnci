package evaluator

import (
	"fmt"
	"mwnci/object"
)

func Println(args ...object.Object) object.Object {
	putsFun(args...)
	fmt.Print("\n")
	return NULL
}
