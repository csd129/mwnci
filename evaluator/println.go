package evaluator

import (
	"fmt"
	"mwnci/object"
)

// Println ...
func Println(args ...object.Object) object.Object {
	putsFun(args...)
	fmt.Print("\n")
	return NULL
}

