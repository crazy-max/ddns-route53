package wanip

import (
	"net"

	"github.com/pkg/errors"
)

var errNoSuitableAddress = errors.New("no suitable address found for interface")

func isUnavailableError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, errNoSuitableAddress) || isUnavailableErrno(err) {
		return true
	}
	var dnsErr *net.DNSError
	return errors.As(err, &dnsErr) && isUnavailableDNSError(dnsErr)
}
