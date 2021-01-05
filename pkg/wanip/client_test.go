package wanip

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	c *Client
)

func TestMain(m *testing.M) {
	c = NewClient(
		fmt.Sprintf("ddns-route53/test go/%s %s", runtime.Version()[2:], strings.Title(runtime.GOOS)),
		3,
	)
	os.Exit(m.Run())
}

func TestClient_IPv4(t *testing.T) {
	ip, errs := c.IPv4()
	if errs != nil && isNetworkUnreachable(errs) {
		t.Skip("Skipping unsupported IPv4 on host")
	}
	assert.NotEmpty(t, ip)
	fmt.Println("IPv4:", ip)
}

func TestClient_IPv6(t *testing.T) {
	ip, errs := c.IPv6()
	if errs != nil && isNetworkUnreachable(errs) {
		t.Skip("Skipping unsupported IPv6 on host")
	}
	assert.NotEmpty(t, ip)
	fmt.Println("IPv6:", ip)
}

func isNetworkUnreachable(errs Errors) bool {
	for _, err := range errs {
		if !(strings.Contains(err.Err.Error(), "no such host") ||
			strings.Contains(err.Err.Error(), "network is unreachable")) {
			return false
		}
	}
	return true
}
