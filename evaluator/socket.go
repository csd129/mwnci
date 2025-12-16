package evaluator

import (
	"strings"
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

// Socket ...
func Socket(args ...object.Object) object.Object {
	if err := typing.Check(
		"socket", args,
		typing.ExactArgs(1),
		typing.WithTypes(object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	var (
		domain int
		typ    int
		proto  int
	)

	arg := args[0].(*object.String).Value

	switch strings.ToLower(arg) {
	case "unix":
		domain = syscall.AF_UNIX
		typ = syscall.SOCK_STREAM
		proto = 0
	case "tcp4":
		domain = syscall.AF_INET
		typ = syscall.SOCK_STREAM
		proto = syscall.IPPROTO_TCP
	case "tcp6":
		domain = syscall.AF_INET6
		typ = syscall.SOCK_STREAM
		proto = syscall.IPPROTO_TCP
	case "udp4":
		domain = syscall.AF_INET
		typ = syscall.SOCK_DGRAM
		proto = syscall.IPPROTO_UDP
	case "udp6":
		domain = syscall.AF_INET6
		typ = syscall.SOCK_DGRAM
		proto = syscall.IPPROTO_UDP
	default:
		return newError("ValueError: invalid socket type '%s'", arg)
	}

	fd, err := syscall.Socket(domain, typ, proto)
	if err != nil {
		return newError("SocketError: %s", err)
	}

	if domain == syscall.AF_INET || domain == syscall.AF_INET6 {
		if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
			return newError("SocketError: cannot enable SO_REUSEADDR: %s", err)
		}
	}

	return &object.Integer{Value: int64(fd)}
}

