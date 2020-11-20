package ipify

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/crazy-max/ddns-route53/v2/pkg/ip_provider/commons"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
)

type Client struct {
	UserAgent  string
	MaxRetries int
}

// NewClient initializes a new ipify client
func NewClient(userAgent string, maxRetries int) (c commons.Client) {
	return &Client{
		UserAgent:  userAgent,
		MaxRetries: maxRetries,
	}
}

// IPv4 returns your IPv4 address from ipify.org service
func (c *Client) IPv4() (net.IP, error) {
	ip, err := c.wanIP("https://api.ipify.org/")
	if err != nil {
		return nil, err
	}
	if ip == nil || !strings.Contains(ip.String(), ".") {
		return nil, fmt.Errorf("invalid IPv4 address: %s", ip.String())
	}
	return ip, nil
}

// IPv6 returns your IPv6 address from ipify.org service
func (c *Client) IPv6() (net.IP, error) {
	ip, err := c.wanIP("https://api64.ipify.org/")
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
			return nil, fmt.Errorf("received invalid status code %d from ipify.org: %s", res.StatusCode, res.Body)
		}

		return net.ParseIP(string(ip)), nil
	}

	return nil, errors.Wrap(err, "request failed")
}
