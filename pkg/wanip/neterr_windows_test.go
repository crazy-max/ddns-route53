//go:build windows

package wanip

import (
	"net"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/windows"
)

func TestIsUnavailableErrno(t *testing.T) {
	t.Parallel()

	assert.True(t, isUnavailableErrno(errors.Wrap(windows.WSAENETUNREACH, "wrapped")))
	assert.True(t, isUnavailableErrno(errors.Wrap(windows.WSAEADDRNOTAVAIL, "wrapped")))
	assert.False(t, isUnavailableErrno(errors.Wrap(windows.WSAEINVAL, "wrapped")))
}

func TestIsUnavailableDNSError(t *testing.T) {
	t.Parallel()

	assert.True(t, isUnavailableDNSError(&net.DNSError{IsNotFound: true}))
	assert.True(t, isUnavailableDNSError(&net.DNSError{Err: "getaddrinfow: " + windows.WSANO_DATA.Error()}))
	assert.True(t, isUnavailableDNSError(&net.DNSError{Err: "lookup foo: " + windows.WSAHOST_NOT_FOUND.Error()}))
	assert.False(t, isUnavailableDNSError(&net.DNSError{Err: "some other dns failure"}))
	assert.False(t, isUnavailableDNSError(nil))
}
