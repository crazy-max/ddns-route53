package wanip

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
)

// Client represents an active wanip object
type Client struct {
	ifname     string
	userAgent  string
	maxRetries int
}

// Option provides a variadic option for configuring the client.
type Option func(b *Client)

// WithInterfaceName sets interface name.
func WithInterfaceName(ifname string) Option {
	return func(b *Client) {
		b.ifname = ifname
	}
}

// WithUserAgent sets user agent.
func WithUserAgent(userAgent string) Option {
	return func(b *Client) {
		b.userAgent = userAgent
	}
}

// WithMaxRetries sets max retries.
func WithMaxRetries(maxRetries int) Option {
	return func(b *Client) {
		b.maxRetries = maxRetries
	}
}

// New initializes a new wanip client
func New(opts ...Option) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// IPv4 returns your IPv4 address
func (c *Client) IPv4() (net.IP, error) {
	return c.lookup(defaultIPv4Providers, false)
}

// IPv6 returns your IPv6 address
func (c *Client) IPv6() (net.IP, error) {
	return c.lookup(defaultIPv6Providers, true)
}

func (c *Client) lookup(providers []provider, wantIPv6 bool) (net.IP, error) {
	httpc, err := c.httpClient(wantIPv6)
	if err != nil {
		return nil, err
	}

	failures := make([]ProviderFailure, 0, len(providers))
	for _, p := range providers {
		ip, err := c.getIP(httpc, p)
		if err != nil {
			failures = append(failures, ProviderFailure{
				URL: p.URL,
				Err: err,
			})
			continue
		}
		if ip != nil && ((wantIPv6 && ip.To16() != nil && ip.To4() == nil) || (!wantIPv6 && ip.To4() != nil)) {
			return ip, nil
		}
		failures = append(failures, ProviderFailure{
			URL: p.URL,
			Err: errors.Errorf("invalid IP address: %s", ip.String()),
		})
	}
	return nil, &ProviderError{Failures: failures}
}

func (c *Client) getIP(httpc *http.Client, p provider) (net.IP, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", p.URL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	res, err := httpc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("received invalid status code %d from %s: %s", res.StatusCode, p.URL, strings.TrimSpace(string(body)))
	}

	ip := p.Parse(body)
	if ip == nil {
		return nil, errors.Errorf("failed to parse IP address from %s", p.URL)
	}
	return ip, nil
}

func (c *Client) httpClient(wantIPv6 bool) (*http.Client, error) {
	rc := retryablehttp.NewClient()
	rc.RetryMax = c.maxRetries
	rc.Logger = nil
	t, err := c.transport(rc.HTTPClient.Transport, wantIPv6)
	if err != nil {
		return nil, err
	}
	rc.HTTPClient.Transport = t
	return rc.StandardClient(), nil
}

func (c *Client) transport(base http.RoundTripper, wantIPv6 bool) (*http.Transport, error) {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	if c.ifname != "" {
		ifadrr, err := interfaceAddress(c.ifname, wantIPv6)
		if err != nil {
			return nil, err
		}
		dialer.LocalAddr = &net.TCPAddr{IP: ifadrr}
	}

	network := "tcp4"
	if wantIPv6 {
		network = "tcp6"
	}

	baseTransport, ok := base.(*http.Transport)
	if !ok || baseTransport == nil {
		baseTransport = http.DefaultTransport.(*http.Transport)
	}

	t := baseTransport.Clone()
	t.DialContext = func(ctx context.Context, _, addr string) (net.Conn, error) {
		return dialer.DialContext(ctx, network, addr)
	}
	return t, nil
}

func interfaceAddress(interfaceName string, wantIPv6 bool) (net.IP, error) {
	addrs, err := interfaceAddresses(interfaceName)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if wantIPv6 {
				if ipnet.IP.To16() != nil && ipnet.IP.To4() == nil {
					return ipnet.IP, nil
				}
			} else {
				if ipnet.IP.To16() != nil && ipnet.IP.To4() != nil {
					return ipnet.IP, nil
				}
			}
		}
	}
	return nil, errors.Wrapf(errNoSuitableAddress, "%s", interfaceName)
}

func interfaceAddresses(interfaceName string) ([]net.Addr, error) {
	if interfaceName == "any" {
		return net.InterfaceAddrs()
	}
	ief, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get interface: %s", interfaceName)
	}
	addrs, err := ief.Addrs()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get interface addresses for: %s", interfaceName)
	}
	return addrs, nil
}
