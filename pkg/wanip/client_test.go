package wanip

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ua = fmt.Sprintf("ddns-route53/test go/%s %s", runtime.Version()[2:], runtime.GOOS)

func TestWanIP(t *testing.T) {
	c := New(WithMaxRetries(3), WithUserAgent(ua))
	testLiveLookup(t, c)
}

func TestCustomInterface(t *testing.T) {
	ifname, err := defaultIfname()
	require.NoError(t, err)

	c := New(WithInterfaceName(ifname), WithMaxRetries(3), WithUserAgent(ua))
	testLiveLookup(t, c)
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

func TestLookupRetriesProviderRequests(t *testing.T) {
	t.Parallel()

	var attempts atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if attempts.Add(1) < 3 {
			http.Error(w, "nope", http.StatusBadGateway)
			return
		}
		_, _ = io.WriteString(w, "203.0.113.42")
	}))
	defer srv.Close()

	c := New(WithMaxRetries(2))
	ip, err := c.lookup([]provider{
		{URL: srv.URL, Parse: parsePlainTextIP},
	}, false)
	require.NoError(t, err)
	require.NotNil(t, ip)
	assert.Equal(t, "203.0.113.42", ip.String())
	assert.EqualValues(t, 3, attempts.Load())
}

func TestLookupHonorsContextCancellation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	started := make(chan struct{}, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case started <- struct{}{}:
		default:
		}
		<-r.Context().Done()
	}))
	defer srv.Close()

	c := New(
		WithContext(ctx),
		WithIPv4Providers([]string{srv.URL}),
	)

	done := make(chan error, 1)
	go func() {
		_, err := c.IPv4()
		done <- err
	}()

	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("lookup did not start request")
	}

	cancel(context.Canceled)

	select {
	case err := <-done:
		require.Error(t, err)
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(time.Second):
		t.Fatal("lookup did not stop after context cancellation")
	}
}

func TestCustomIPv4ProvidersReplaceDefaults(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "203.0.113.42")
	}))
	defer srv.Close()

	c := New(WithIPv4Providers([]string{srv.URL}))
	ip, err := c.IPv4()
	require.NoError(t, err)
	require.NotNil(t, ip)
	assert.Equal(t, "203.0.113.42", ip.String())
}

func TestCustomIPv6ProvidersReplaceDefaults(t *testing.T) {
	t.Parallel()

	listener, err := (&net.ListenConfig{}).Listen(context.Background(), "tcp6", "[::1]:0")
	if err != nil {
		t.Skipf("Skipping IPv6 custom provider test: %v", err)
	}

	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "2001:db8::42")
	}))
	srv.Listener = listener
	srv.Start()
	defer srv.Close()

	c := New(WithIPv6Providers([]string{srv.URL}))
	ip, err := c.IPv6()
	require.NoError(t, err)
	require.NotNil(t, ip)
	assert.Equal(t, "2001:db8::42", ip.String())
}

func testLiveLookup(t *testing.T, c *Client) {
	t.Helper()

	cases := []struct {
		name   string
		family string
		lookup func() (net.IP, error)
	}{
		{
			name:   "v4",
			family: "IPv4",
			lookup: c.IPv4,
		},
		{
			name:   "v6",
			family: "IPv6",
			lookup: c.IPv6,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ip, err := tt.lookup()
			if err != nil && isNetworkUnreachable(err) {
				t.Skipf("Skipping unsupported %s on host", tt.family)
			}
			if ip == nil && err != nil {
				t.Logf("%s errors: %+v", tt.family, providerFailures(err))
			}
			assert.NotEmpty(t, ip)
			t.Logf("%s: %s", tt.family, ip)
		})
	}
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
		if !isUnavailableError(failure.Err) && !hasSandboxNetworkUnreachableMessage(failure.Err) {
			return false
		}
	}
	return true
}

var sandboxNetworkUnreachablePatterns = []string{
	"forbidden by its access permissions",
	"socket operation was attempted to an unreachable network",
}

func hasSandboxNetworkUnreachableMessage(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	for _, pattern := range sandboxNetworkUnreachablePatterns {
		if strings.Contains(msg, pattern) {
			return true
		}
	}
	return false
}
