package evaluator

import ( 
	"mwnci/object"
	"mwnci/typing"
	"syscall"
)

func Fgetc(args ...object.Object) object.Object {
	if err := typing.Check(
		"fgetc", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	var (
		fd int
		n  = 1
	)

	fd = int(args[0].(*object.Integer).Value)

	if len(args) == 2 {
		n = int(args[1].(*object.Integer).Value)
	}

	buf := make([]byte, n)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		return newError("IOError: %s", err)
	}
	charbyte := string(buf[:n])
	return &object.String{Value: charbyte}
}
