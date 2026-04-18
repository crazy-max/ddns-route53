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
	ctx        context.Context
	hc         *http.Client
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
	c := &Client{
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(c)
	}

	rc := retryablehttp.NewClient()
	rc.RetryMax = c.maxRetries
	rc.Logger = nil
	c.hc = rc.StandardClient()
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

func (c *Client) lookup(providers []provider, v6 bool) (net.IP, error) {
	var failures []ProviderFailure
	for _, p := range providers {
		ip, err := c.getIP(p, v6)
		if err != nil {
			failures = append(failures, ProviderFailure{
				URL: p.URL,
				Err: err,
			})
			continue
		}
		if ip != nil && ((v6 && ip.To16() != nil && ip.To4() == nil) || (!v6 && ip.To4() != nil)) {
			return ip, nil
		}
		failures = append(failures, ProviderFailure{
			URL: p.URL,
			Err: errors.Errorf("invalid IP address: %s", ip.String()),
		})
	}
	return nil, &ProviderError{Failures: failures}
}

func (c *Client) getIP(p provider, v6 bool) (net.IP, error) {
	httpc := *c.hc
	if t, err := c.transport(v6); err != nil {
		return nil, err
	} else {
		httpc.Transport = t
	}

	req, err := http.NewRequestWithContext(c.ctx, "GET", p.URL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	if c.userAgent != "" {
		req.Header.Add("User-Agent", c.userAgent)
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

// transport returns transport that uses the specified interface.
func (c *Client) transport(v6 bool) (*http.Transport, error) {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	if c.ifname != "" {
		ifadrr, err := interfaceAddress(c.ifname, v6)
		if err != nil {
			return nil, err
		}
		dialer.LocalAddr = &net.TCPAddr{IP: ifadrr}
	}

	// same as http.DefaultTransport but with a custom dialer with local addr
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
			network := "tcp4"
			if v6 {
				network = "tcp6"
			}
			return dialer.DialContext(ctx, network, addr)
		},
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}, nil
}

// interfaceAddress returns the first IPv4/IPv6 address of the given interface.
func interfaceAddress(interfaceName string, v6 bool) (net.IP, error) {
	addrs, err := interfaceAddresses(interfaceName)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if v6 {
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

// interfaceAddresses returns all interface addresses.
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
