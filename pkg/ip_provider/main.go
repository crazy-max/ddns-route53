package ip_provider

import (
	"github.com/crazy-max/ddns-route53/v2/pkg/ip_provider/commons"
	"github.com/crazy-max/ddns-route53/v2/pkg/ip_provider/providers/identme"
	"github.com/crazy-max/ddns-route53/v2/pkg/ip_provider/providers/ipify"
)

type Client = commons.Client

// NewClient initializes a new ipify client
func NewClient(provider string, userAgent string, maxRetries int) (c Client) {
	if provider == "ipify" {
		return ipify.NewClient(userAgent, maxRetries)
	}

	return identme.NewClient(userAgent, maxRetries)
}
