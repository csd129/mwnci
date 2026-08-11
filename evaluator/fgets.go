package evaluator

import ( 
	"mwnci/object"
	"mwnci/typing"
	"strings"
	"syscall"
)

func Fgets(args ...object.Object) object.Object {
	if err := typing.Check(
		"fgets", args,
		typing.RangeOfArgs(1,2),
		typing.WithTypes(object.INTEGER_OBJ, object.INTEGER_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	var (
		fd int
		n  = DefaultBufferSize
	)

	fd = int(args[0].(*object.Integer).Value)
	if len(args) == 2 {
		n = int(args[1].(*object.Integer).Value)
	}

	fullstring := ""
	buf := make([]byte, n)
        offset, err := syscall.Seek(fd, 0, 1)
	if err != nil {
		return newError("IOError: %s", err)
	}
	n, err = syscall.Read(fd, buf)
	if len(buf[:n]) == 0 {return FALSE}
	if err != nil {
		return newError("IOError: %s", err)
	}
	index := strings.Index(string(buf), "\n")
	if index == -1 {
		fullstring=string(buf)
	} else {
		fullstring=string(buf[:index])
	}
	offset=offset + int64(len(fullstring)) + 1
	_, err = syscall.Seek(fd, offset, 0)
	if err != nil {
		return newError("IOError: %s", err)
	}
	return &object.String{Value: fullstring}
}
