package wanip

import (
	"net"
	"strings"
)

type ProviderFailure struct {
	URL string
	Err error
}

type ProviderError struct {
	Failures []ProviderFailure
}

func (e *ProviderError) Error() string {
	if e == nil {
		return ""
	}
	return "all WAN IP providers failed"
}

func (e *ProviderError) Unwrap() []error {
	if e == nil {
		return nil
	}
	errs := make([]error, 0, len(e.Failures))
	for _, failure := range e.Failures {
		errs = append(errs, failure.Err)
	}
	return errs
}

type provider struct {
	URL   string
	Parse func([]byte) net.IP
}

var defaultIPv4Providers = []provider{
	{URL: "https://checkip.global.api.aws", Parse: parsePlainTextIP},
	{URL: "https://checkip.amazonaws.com", Parse: parsePlainTextIP},
	{URL: "https://cloudflare.com/cdn-cgi/trace", Parse: parseCloudflareTraceIP},
	{URL: "https://ipv4.nsupdate.info/myip", Parse: parsePlainTextIP},
}

var defaultIPv6Providers = []provider{
	{URL: "https://checkip.global.api.aws", Parse: parsePlainTextIP},
	{URL: "https://cloudflare.com/cdn-cgi/trace", Parse: parseCloudflareTraceIP},
	{URL: "https://ipv6.nsupdate.info/myip", Parse: parsePlainTextIP},
}

func parsePlainTextIP(body []byte) net.IP {
	return net.ParseIP(strings.TrimSpace(string(body)))
}

func parseCloudflareTraceIP(body []byte) net.IP {
	for line := range strings.SplitSeq(string(body), "\n") {
		if rest, ok := strings.CutPrefix(strings.TrimSpace(line), "ip="); ok {
			return net.ParseIP(strings.TrimSpace(rest))
		}
	}
	return nil
}

func plainTextProviders(urls []string) []provider {
	providers := make([]provider, 0, len(urls))
	for _, raw := range urls {
		providers = append(providers, provider{
			URL:   raw,
			Parse: parsePlainTextIP,
		})
	}
	return providers
}
