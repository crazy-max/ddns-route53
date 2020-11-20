package commons

import "net"

type Client interface {
	IPv4() (net.IP, error)
	IPv6() (net.IP, error)
}
