//go:build !windows

package wanip

import (
	"net"
	"syscall"

	"github.com/pkg/errors"
)

func isUnavailableErrno(err error) bool {
	return errors.Is(err, syscall.ENETUNREACH) ||
		errors.Is(err, syscall.EHOSTUNREACH) ||
		errors.Is(err, syscall.EADDRNOTAVAIL) ||
		errors.Is(err, syscall.EAFNOSUPPORT)
}

func isUnavailableDNSError(err *net.DNSError) bool {
	return err != nil && err.IsNotFound
}
