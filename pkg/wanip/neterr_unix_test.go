//go:build !windows

package wanip

import (
	"net"
	"syscall"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestIsUnavailableErrno(t *testing.T) {
	t.Parallel()

	assert.True(t, isUnavailableErrno(errors.Wrap(syscall.ENETUNREACH, "wrapped")))
	assert.True(t, isUnavailableErrno(errors.Wrap(syscall.EADDRNOTAVAIL, "wrapped")))
	assert.False(t, isUnavailableErrno(errors.Wrap(syscall.EINVAL, "wrapped")))
}

func TestIsUnavailableDNSError(t *testing.T) {
	t.Parallel()

	assert.True(t, isUnavailableDNSError(&net.DNSError{IsNotFound: true}))
	assert.False(t, isUnavailableDNSError(&net.DNSError{IsNotFound: false}))
	assert.False(t, isUnavailableDNSError(nil))
}
