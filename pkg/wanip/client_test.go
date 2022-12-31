package wanip

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"strings"
	"testing"

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
				ip, errs := c.IPv6()
				if errs != nil && isNetworkUnreachable(errs) {
					t.Skip("Skipping unsupported IPv6 on host")
				}
				assert.NotEmpty(t, ip)
				fmt.Println("IPv6:", ip)
			} else {
				ip, errs := c.IPv4()
				if errs != nil && isNetworkUnreachable(errs) {
					t.Skip("Skipping unsupported IPv4 on host")
				}
				assert.NotEmpty(t, ip)
				fmt.Println("IPv4:", ip)
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
				ip, errs := c.IPv6()
				if errs != nil && isNetworkUnreachable(errs) {
					t.Skip("Skipping unsupported IPv6 on host")
				}
				assert.NotEmpty(t, ip)
				fmt.Println("IPv6:", ip)
			} else {
				ip, errs := c.IPv4()
				if errs != nil && isNetworkUnreachable(errs) {
					t.Skip("Skipping unsupported IPv4 on host")
				}
				assert.NotEmpty(t, ip)
				fmt.Println("IPv4:", ip)
			}
		})
	}
}

func isNetworkUnreachable(errs Errors) bool {
	for _, err := range errs {
		if !(strings.Contains(err.Err.Error(), "no such host") ||
			strings.Contains(err.Err.Error(), "network is unreachable") ||
			strings.Contains(err.Err.Error(), "cannot assign requested address") ||
			strings.Contains(err.Err.Error(), "no suitable address found for interface")) {
			return false
		}
	}
	return true
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
