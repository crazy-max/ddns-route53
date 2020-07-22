package identme

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

// Client represents an active identme object
type Client struct {
	UserAgent  string
	MaxRetries int
}

// NewClient initializes a new identme client
func NewClient(userAgent string, maxRetries int) (c *Client) {
	return &Client{
		UserAgent:  userAgent,
		MaxRetries: maxRetries,
	}
}

// IPv4 returns your IPv4 address from ident.me service
func (c *Client) IPv4() (net.IP, error) {
	ip, err := c.wanIP("https://v4.ident.me/")
	if err != nil {
		return nil, err
	}
	if ip == nil || !strings.Contains(ip.String(), ".") {
		return nil, fmt.Errorf("invalid IPv4 address: %s", ip.String())
	}
	return ip, nil
}

// IPv6 returns your IPv6 address from ident.me service
func (c *Client) IPv6() (net.IP, error) {
	ip, err := c.wanIP("https://v6.ident.me/")
	if err != nil {
		return nil, err
	}
	if ip == nil || !strings.Contains(ip.String(), ":") {
		return nil, fmt.Errorf("invalid IPv6 address: %s", ip.String())
	}
	return ip, nil
}

func (c *Client) wanIP(apiURL string) (net.IP, error) {
	var err error
	var req *http.Request
	var res *http.Response

	client := &http.Client{}
	b := &backoff.Backoff{
		Jitter: true,
	}

	req, err = http.NewRequest("GET", apiURL, nil)
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
			return nil, fmt.Errorf("received invalid status code %d from ident.me: %s", res.StatusCode, res.Body)
		}

		return net.ParseIP(string(ip)), nil
	}

	return nil, errors.Wrap(err, "request failed")
}
