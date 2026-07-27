package evaluator

import ( 
	"mwnci/object"
	"mwnci/typing"
	"syscall"
)

func Fgets(args ...object.Object) object.Object {
	if err := typing.Check(
		"fgets", args,
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
	fullstring := ""
	buf := make([]byte, n)
        for string(buf[:n]) != "\n" {
		n, err := syscall.Read(fd, buf)
                if len(buf[:n]) == 0 {return FALSE}
		if err != nil {
			return newError("IOError: %s", err)
		}
		charbyte := string(buf[:n])
		fullstring = fullstring + charbyte
        }
	return &object.String{Value: fullstring}
}
