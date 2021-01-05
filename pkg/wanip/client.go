package wanip

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/jpillora/backoff"
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
	UserAgent  string
	MaxRetries int
}

// NewClient initializes a new wanip client
func NewClient(userAgent string, maxRetries int) (c *Client) {
	return &Client{
		UserAgent:  userAgent,
		MaxRetries: maxRetries,
	}
}

// IPv4 returns your IPv4 address
func (c *Client) IPv4() (net.IP, Errors) {
	var errs Errors
	for _, providerURL := range []string{
		"https://ipv4.nsupdate.info/myip",
		"http://ipv4.icanhazip.com",
		"https://v4.ident.me",
		"http://ipv4.yunohost.org",
		"https://ipv4.wtfismyip.com/text",
	} {
		ip, err := c.getIP(providerURL)
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
		"http://ipv6.icanhazip.com",
		"https://v6.ident.me",
		"http://ipv6.yunohost.org",
		"https://ipv6.wtfismyip.com/text",
	} {
		ip, err := c.getIP(providerURL)
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

func (c *Client) getIP(providerURL string) (net.IP, error) {
	var err error
	var req *http.Request
	var res *http.Response

	client := &http.Client{}
	b := &backoff.Backoff{
		Jitter: true,
	}

	req, err = http.NewRequest("GET", providerURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	req.Header.Add("User-Agent", c.UserAgent)
	for tries := 0; tries < c.MaxRetries; tries++ {
		res, err = client.Do(req)
		if err != nil {
			time.Sleep(b.Duration())
			continue
		}
		defer res.Body.Close()

		ip, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrap(err, "request failed")
		}

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("received invalid status code %d from %s: %s", res.StatusCode, providerURL, res.Body)
		}

		return net.ParseIP(string(ip)), nil
	}

	return nil, errors.Wrap(err, "request failed")
}
