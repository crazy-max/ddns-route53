package wanip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ua = fmt.Sprintf("ddns-route53/test go/%s %s", runtime.Version()[2:], runtime.GOOS)

func TestWanIP(t *testing.T) {
	c := New(WithMaxRetries(3), WithUserAgent(ua))

	cases := []struct {
		name string
		v6   bool
	}{
		{
			name: "v4",
			v6:   false,
		},
		{
			name: "v6",
			v6:   true,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.v6 {
				ip, err := c.IPv6()
				if ip == nil && err == nil {
					t.Skip("Skipping unsupported IPv6 on host")
				}
				if err != nil && isNetworkUnreachable(err) {
					t.Skip("Skipping unsupported IPv6 on host")
				}
				if ip == nil && err != nil {
					t.Logf("IPv6 errors: %+v", providerFailures(err))
				}
				assert.NotEmpty(t, ip)
				t.Logf("IPv6: %s", ip)
			} else {
				ip, err := c.IPv4()
				if err != nil && isNetworkUnreachable(err) {
					t.Skip("Skipping unsupported IPv4 on host")
				}
				if ip == nil && err != nil {
					t.Logf("IPv4 errors: %+v", providerFailures(err))
				}
				assert.NotEmpty(t, ip)
				t.Logf("IPv4: %s", ip)
			}
		})
	}
}

func TestCustomInterface(t *testing.T) {
	ifname, err := defaultIfname()
	require.NoError(t, err)

	c := New(WithInterfaceName(ifname), WithMaxRetries(3), WithUserAgent(ua))

	cases := []struct {
		name string
		v6   bool
	}{
		{
			name: "v4",
			v6:   false,
		},
		{
			name: "v6",
			v6:   true,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.v6 {
				ip, err := c.IPv6()
				if ip == nil && err == nil {
					t.Skip("Skipping unsupported IPv6 on host")
				}
				if err != nil && isNetworkUnreachable(err) {
					t.Skip("Skipping unsupported IPv6 on host")
				}
				if ip == nil && err != nil {
					t.Logf("IPv6 errors: %+v", providerFailures(err))
				}
				assert.NotEmpty(t, ip)
				t.Logf("IPv6: %s", ip)
			} else {
				ip, err := c.IPv4()
				if err != nil && isNetworkUnreachable(err) {
					t.Skip("Skipping unsupported IPv4 on host")
				}
				if ip == nil && err != nil {
					t.Logf("IPv4 errors: %+v", providerFailures(err))
				}
				assert.NotEmpty(t, ip)
				t.Logf("IPv4: %s", ip)
			}
		})
	}
}

func TestIPv6UnavailableReturnsProviderError(t *testing.T) {
	t.Parallel()

	c := New()
	ip, err := c.lookup([]provider{
		{URL: "https://example.invalid", Parse: parsePlainTextIP},
		{URL: "https://example.invalid", Parse: parseCloudflareTraceIP},
	}, true)

	assert.Nil(t, ip)
	require.Error(t, err)

	var providerErr *ProviderError
	require.True(t, errors.As(err, &providerErr))
	require.Len(t, providerErr.Failures, 2)
	for _, failure := range providerErr.Failures {
		assert.True(t, isUnavailableError(failure.Err))
	}
}

func TestParsePlainTextIP(t *testing.T) {
	t.Parallel()

	ip := parsePlainTextIP([]byte(" 203.0.113.10 \n"))
	require.NotNil(t, ip)
	assert.Equal(t, "203.0.113.10", ip.String())
}

func TestParseCloudflareTraceIP(t *testing.T) {
	t.Parallel()

	ip := parseCloudflareTraceIP([]byte("fl=29f5\nip=2001:db8::1\nts=1\n"))
	require.NotNil(t, ip)
	assert.Equal(t, "2001:db8::1", ip.String())
}

func TestIPv4FallsBackToCloudflareTrace(t *testing.T) {
	t.Parallel()

	awsGlobal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "nope", http.StatusBadGateway)
	}))
	defer awsGlobal.Close()

	awsLegacy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "still not an ip")
	}))
	defer awsLegacy.Close()

	cloudflare := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "fl=29f5\nip=203.0.113.42\nts=1\n")
	}))
	defer cloudflare.Close()

	nsupdate := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		t.Fatal("nsupdate fallback should not be used after cloudflare succeeds")
	}))
	defer nsupdate.Close()

	c := New()
	ip, err := c.lookup([]provider{
		{URL: awsGlobal.URL, Parse: parsePlainTextIP},
		{URL: awsLegacy.URL, Parse: parsePlainTextIP},
		{URL: cloudflare.URL, Parse: parseCloudflareTraceIP},
		{URL: nsupdate.URL, Parse: parsePlainTextIP},
	}, false)
	require.NoError(t, err)
	require.NotNil(t, ip)
	assert.Equal(t, "203.0.113.42", ip.String())
}

func defaultIfname() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To16() != nil || ipnet.IP.To4() != nil {
					return i.Name, nil
				}
			}
		}
	}
	return "", errors.New("no default interface found")
}

func providerFailures(err error) []ProviderFailure {
	if err == nil {
		return nil
	}
	var providerErr *ProviderError
	if errors.As(err, &providerErr) {
		return providerErr.Failures
	}
	return []ProviderFailure{{Err: err}}
}

func isNetworkUnreachable(err error) bool {
	for _, failure := range providerFailures(err) {
		if !isUnavailableError(failure.Err) &&
			!strings.Contains(failure.Err.Error(), "forbidden by its access permissions") &&
			!strings.Contains(failure.Err.Error(), "socket operation was attempted to an unreachable network") {
			return false
		}
	}
	return true
}
