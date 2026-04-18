package wanip

import (
	"net"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestIsUnavailableError(t *testing.T) {
	t.Parallel()

	assert.True(t, isUnavailableError(errNoSuitableAddress))
	assert.True(t, isUnavailableError(errors.Wrap(&net.DNSError{IsNotFound: true}, "wrapped")))
	assert.False(t, isUnavailableError(errors.New("nope")))
}
