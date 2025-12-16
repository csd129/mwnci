package evaluator

import (
	"syscall"

	"mwnci/object"
	"mwnci/typing"
)

// Connect ...
func Connect(args ...object.Object) object.Object {
	if err := typing.Check(
		"connect", args,
		typing.ExactArgs(2),
		typing.WithTypes(object.INTEGER_OBJ, object.STRING_OBJ),
	); err != nil {
		return newError("%s", err.Error())
	}

	var sa syscall.Sockaddr

	fd := int(args[0].(*object.Integer).Value)
	address := args[1].(*object.String).Value

	sockaddr, err := syscall.Getsockname(fd)
	if err != nil {
		return newError("ValueError: %s", err)
	}

	if _, ok := sockaddr.(*syscall.SockaddrInet4); ok {
		addr, port, err := parseV4Address(address)
		if err != nil {
			return newError("ValueError: Invalid IPv4 address '%s': %s", address, err)
		}
		sa = &syscall.SockaddrInet4{Addr: addr, Port: port}
	} else if _, ok := sockaddr.(*syscall.SockaddrInet6); ok {
		addr, port, err := parseV6Address(address)
		if err != nil {
			return newError("ValueError: Invalid IPv6 address '%s': %s", address, err)
		}
		sa = &syscall.SockaddrInet6{Addr: addr, Port: port}
	} else {
		return newError("ValueError: Invalid socket type %T for bind '%s'", sockaddr, address)
	}

	if err = syscall.Connect(fd, sa); err != nil {
		return newError("SocketError: host %s, %s", address, err)
	}

	return NULL
}
