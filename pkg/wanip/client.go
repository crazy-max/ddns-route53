package wanip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
)

// Errors holds slice of wanip errors
type Errors []Error

// Error holds wanip error
type Error struct {
	Err         error
	ProviderURL string
}

// Client represents an active wanip object
type Client struct {
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
	c := &Client{}
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
func (c *Client) IPv4() (net.IP, Errors) {
	var errs Errors
	for _, providerURL := range []string{
		"https://ipv4.nsupdate.info/myip",
		"https://v4.ident.me",
		"https://ipv4.yunohost.org",
		"https://ipv4.wtfismyip.com/text",
	} {
		ip, err := c.getIP(providerURL, false)
		if err != nil {
			errs = append(errs, Error{
				Err:         err,
				ProviderURL: providerURL,
			})
			continue
		}
		if ip != nil && len(ip.To4()) == net.IPv4len && strings.Contains(ip.String(), ".") {
			return ip, nil
		}
		errs = append(errs, Error{
			Err:         fmt.Errorf("invalid IPv4 address: %s", ip.String()),
			ProviderURL: providerURL,
		})
	}
	return nil, errs
}

// IPv6 returns your IPv6 address
func (c *Client) IPv6() (net.IP, Errors) {
	var errs Errors
	for _, providerURL := range []string{
		"https://ipv6.nsupdate.info/myip",
		"https://v6.ident.me",
		"https://ipv6.yunohost.org",
		"https://ipv6.wtfismyip.com/text",
	} {
		ip, err := c.getIP(providerURL, true)
		if err != nil {
			errs = append(errs, Error{
				Err:         err,
				ProviderURL: providerURL,
			})
			continue
		}
		if ip != nil && len(ip.To16()) == net.IPv6len && strings.Contains(ip.String(), ":") {
			return ip, nil
		}
		errs = append(errs, Error{
			Err:         fmt.Errorf("invalid IPv6 address: %s", ip.String()),
			ProviderURL: providerURL,
		})
	}
	return nil, errs
}

func (c *Client) getIP(providerURL string, v6 bool) (net.IP, error) {
	httpc := c.hc
	if t, err := c.transport(v6); err != nil {
		return nil, err
	} else if t != nil {
		httpc.Transport = t
	}

	req, err := http.NewRequest("GET", providerURL, nil)
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

	ip, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received invalid status code %d from %s: %s", res.StatusCode, providerURL, res.Body)
	}

	return net.ParseIP(string(ip)), nil
}

// transport returns transport that uses the specified interface.
func (c *Client) transport(v6 bool) (*http.Transport, error) {
	if c.ifname == "" {
		return nil, nil
	}

	ifadrr, err := interfaceAddress(c.ifname, v6)
	if err != nil {
		return nil, err
	}

	// same as http.DefaultTransport but with a custom dialer with local addr
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: &net.TCPAddr{IP: ifadrr},
		}).DialContext,
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
	return nil, fmt.Errorf("no suitable address found for interface: %s", interfaceName)
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
