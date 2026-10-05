package app

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/crazy-max/ddns-route53/v2/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunSRV(t *testing.T) {
	cases := []struct {
		name         string
		currentValue string
		currentTTL   int
		mixed        bool
		wanFailure   bool
		wantChanges  int
	}{
		{name: "create without WAN IP", wantChanges: 1},
		{name: "unchanged", currentValue: "0 0 2032 host1.example.org.", currentTTL: 300},
		{name: "changed priority", currentValue: "10 0 2032 host1.example.org.", currentTTL: 300, wantChanges: 1},
		{name: "changed weight", currentValue: "0 10 2032 host1.example.org.", currentTTL: 300, wantChanges: 1},
		{name: "changed port", currentValue: "0 0 22 host1.example.org.", currentTTL: 300, wantChanges: 1},
		{name: "changed target", currentValue: "0 0 2032 host2.example.org.", currentTTL: 300, wantChanges: 1},
		{name: "changed TTL", currentValue: "0 0 2032 host1.example.org.", currentTTL: 600, wantChanges: 1},
		{name: "mixed A AAAA and SRV", mixed: true, wantChanges: 3},
		{name: "WAN lookup failure does not block SRV", mixed: true, wanFailure: true, wantChanges: 1},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var updates []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/ipv4" || r.URL.Path == "/ipv6" {
					if tt.wanFailure {
						http.Error(w, "lookup failed", http.StatusBadRequest)
					} else if r.URL.Path == "/ipv4" {
						_, _ = io.WriteString(w, "203.0.113.10")
					} else {
						_, _ = io.WriteString(w, "2001:db8::10")
					}
					return
				}

				assert.Equal(t, "/2013-04-01/hostedzone/ZTEST123/rrset", r.URL.Path)
				w.Header().Set("Content-Type", "application/xml")
				if r.Method == http.MethodGet {
					var records string
					if tt.currentValue != "" {
						records = fmt.Sprintf(`<ResourceRecordSet>
<Name>_ssh._tcp.host1.example.org.</Name><Type>SRV</Type><TTL>%d</TTL>
<ResourceRecords><ResourceRecord><Value>%s</Value></ResourceRecord></ResourceRecords>
</ResourceRecordSet>`, tt.currentTTL, tt.currentValue)
					}
					_, _ = fmt.Fprintf(w, `<ListResourceRecordSetsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
<IsTruncated>false</IsTruncated><MaxItems>100</MaxItems>
<ResourceRecordSets>%s</ResourceRecordSets></ListResourceRecordSetsResponse>`, records)
					return
				}

				assert.Equal(t, http.MethodPost, r.Method)
				body, err := io.ReadAll(r.Body)
				assert.NoError(t, err)
				updates = append(updates, string(body))
				_, _ = io.WriteString(w, `<ChangeResourceRecordSetsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
<ChangeInfo><Id>/change/test</Id><Status>PENDING</Status><SubmittedAt>2026-01-01T00:00:00Z</SubmittedAt></ChangeInfo>
</ChangeResourceRecordSetsResponse>`)
			}))
			defer server.Close()
			t.Setenv("AWS_ENDPOINT_URL_ROUTE_53", server.URL)
			ipv6URL := server.URL
			if tt.mixed {
				listener, err := (&net.ListenConfig{}).Listen(context.Background(), "tcp6", "[::1]:0")
				if err != nil {
					t.Skipf("IPv6 loopback unavailable: %v", err)
				}
				ipv6Server := httptest.NewUnstartedServer(server.Config.Handler)
				_ = ipv6Server.Listener.Close()
				ipv6Server.Listener = listener
				ipv6Server.Start()
				defer ipv6Server.Close()
				ipv6URL = ipv6Server.URL
			}

			cfg := &config.Config{
				Cli: config.Cli{MaxRetries: 1},
				Credentials: &config.Credentials{
					AccessKeyID:     "test-access-key",
					SecretAccessKey: "test-secret-key",
				},
				Route53: &config.Route53{
					HostedZoneID: "ZTEST123",
					HandleIPv4:   new(tt.mixed),
					HandleIPv6:   new(tt.mixed),
					RecordsSet: config.RecordsSet{{
						Name:   "_ssh._tcp.host1.example.org.",
						Type:   "SRV",
						TTL:    300,
						Port:   2032,
						Target: "host1.example.org.",
					}},
				},
				WanIP: &config.WanIP{
					Providers: &config.WanIPProviders{
						IPv4: []string{server.URL + "/ipv4"},
						IPv6: []string{ipv6URL + "/ipv6"},
					},
				},
			}
			if tt.mixed {
				cfg.Route53.RecordsSet = append(cfg.Route53.RecordsSet,
					config.RecordSet{Name: "host1.example.org.", Type: "A", TTL: 300},
					config.RecordSet{Name: "host1.example.org.", Type: "AAAA", TTL: 300},
				)
			}
			app, err := New(context.Background(), model.Meta{}, cfg)
			require.NoError(t, err)
			app.Run()

			if tt.wantChanges == 0 {
				assert.Empty(t, updates)
				return
			}
			require.Len(t, updates, 1)
			assert.Equal(t, tt.wantChanges, strings.Count(updates[0], "<Action>UPSERT</Action>"))
			assert.Contains(t, updates[0], "<Name>_ssh._tcp.host1.example.org.</Name>")
			assert.Contains(t, updates[0], "<Type>SRV</Type>")
			assert.Contains(t, updates[0], "<TTL>300</TTL>")
			assert.Contains(t, updates[0], "<Value>0 0 2032 host1.example.org.</Value>")
			if tt.mixed && !tt.wanFailure {
				assert.Contains(t, updates[0], "<Type>A</Type>")
				assert.Contains(t, updates[0], "<Value>203.0.113.10</Value>")
				assert.Contains(t, updates[0], "<Type>AAAA</Type>")
				assert.Contains(t, updates[0], "<Value>2001:db8::10</Value>")
			}
		})
	}
}
