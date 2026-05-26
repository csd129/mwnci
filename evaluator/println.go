package evaluator

import (
	"fmt"
	"mwnci/object"
)

func Println(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Print(arg)
	}
	fmt.Print("\n")
	return NULL
}
