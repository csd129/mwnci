package evaluator

import (
	"fmt"
	"os"
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

func parseMode(mode string) (int, error) {
	var flag int
	for _, c := range mode {
		switch c {
		case 'r':
			if (flag & os.O_WRONLY) != 0 {
				flag |= os.O_RDWR
			} else {
				flag |= os.O_RDONLY
			}
		case 'w':
			if (flag & os.O_RDONLY) != 0 {
				flag |= os.O_RDWR
			} else {
				flag |= os.O_WRONLY
			}
		case 'a':
			flag |= os.O_APPEND
		default:
			return 0, fmt.Errorf("ValueError: mode string must be one of 'r', 'w', 'a', not '%c'", c)
		}
	}

	if (flag&os.O_WRONLY) != 0 || (flag&os.O_RDWR) != 0 {
		flag |= os.O_CREATE
		if (flag & os.O_APPEND) == 0 {
			flag |= os.O_TRUNC
		}
	}

	if !((flag == os.O_RDONLY) || (flag&os.O_WRONLY != 0) || (flag&os.O_RDWR != 0)) {
		return 0, fmt.Errorf("ValueError: mode string must be at least one of 'r', 'w', or 'rw'")
	}

	return flag, nil
}

// Open ...
func FOpen(args ...object.Object) object.Object {
	if err := typing.Check(
		"fopen", args,
		typing.RangeOfArgs(1, 2),
		typing.WithTypes(object.STRING_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError(err.Error())
	}

	var (
		filename string
		mode            = "r"
		perm     uint32 = 0640
	)

	filename = args[0].(*object.String).Value

	if len(args) == 2 {
		mode = args[1].(*object.String).Value
	}

	flag, err := parseMode(mode)
	if err != nil {
		return newError(err.Error())
	}

	fd, err := syscall.Open(filename, flag, perm)
	if err != nil {
		return newError("IOError: %s", err)
	}

	return &object.Integer{Value: int64(fd)}
}

