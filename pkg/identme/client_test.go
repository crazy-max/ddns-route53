package identme

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	ip, err := c.IPv4()
	if err != nil && isNetworkUnavailable(err) {
		t.Skip("Skipping unsupported IPv4 on host")
	}
	require.NoError(t, err)
	assert.NotEmpty(t, ip)
	fmt.Println("IPv4:", ip)
}

func TestClient_IPv6(t *testing.T) {
	ip, err := c.IPv6()
	if err != nil && isNetworkUnavailable(err) {
		t.Skip("Skipping unsupported IPv6 on host")
	}
	require.NoError(t, err)
	assert.NotEmpty(t, ip)
	fmt.Println("IPv6:", ip)
}

func isNetworkUnavailable(err error) bool {
	return strings.Contains(err.Error(), "no such host") ||
		strings.Contains(err.Error(), "network is unreachable")
}
