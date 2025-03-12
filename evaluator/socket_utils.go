package evaluator

import (
	"fmt"
	"net"
	"strconv"
)

func parseAddress(address string) (ip net.IP, port int, err error) {
	var (
		h string
		p string
		n int64
	)

	h, p, err = net.SplitHostPort(address)
	if err != nil {
		return
	}

	ip = net.ParseIP(h)
	if ip == nil {
		var addrs []string
		addrs, err = net.LookupHost(h)
		if err != nil {
			err = fmt.Errorf("error resolving host '%s'", address)
			return
		}

		if len(addrs) == 0 {
			err = fmt.Errorf("host not found '%s'", address)
			return
		}

		ip = net.ParseIP(addrs[0])
		if ip == nil {
			err = fmt.Errorf("invalid IP address '%s'", address)
			return
		}
	}

	n, err = strconv.ParseInt(p, 10, 16)
	if err != nil {
		return
	}
	port = int(n)
	return
}

func parseV4Address(address string) (addr [4]byte, port int, err error) {
	var ip net.IP
	ip, port, err = parseAddress(address)
	if err != nil {
		return
	}
	copy(addr[:], ip.To4()[:4])
	return
}

func parseV6Address(address string) (addr [16]byte, port int, err error) {
	var ip net.IP
	ip, port, err = parseAddress(address)
	if err != nil {
		return
	}
	copy(addr[:], ip.To16()[:16])
	return
}
