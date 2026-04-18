//go:build windows

package wanip

import (
	"net"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/sys/windows"
)

func isUnavailableErrno(err error) bool {
	return errors.Is(err, windows.WSAENETUNREACH) ||
		errors.Is(err, windows.WSAEHOSTUNREACH) ||
		errors.Is(err, windows.WSAEADDRNOTAVAIL) ||
		errors.Is(err, windows.WSAEAFNOSUPPORT) ||
		errors.Is(err, windows.WSAHOST_NOT_FOUND) ||
		errors.Is(err, windows.WSANO_DATA)
}

func isUnavailableDNSError(err *net.DNSError) bool {
	if err == nil {
		return false
	}
	if err.IsNotFound {
		return true
	}
	return strings.HasSuffix(err.Err, windows.WSAHOST_NOT_FOUND.Error()) ||
		strings.HasSuffix(err.Err, windows.WSANO_DATA.Error())
}
