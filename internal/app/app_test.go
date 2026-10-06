package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/crazy-max/ddns-route53/v2/pkg/route53"
	"github.com/crazy-max/ddns-route53/v2/pkg/wanip"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestRunLogsEachWANProviderFailure(t *testing.T) {
	for _, tc := range []struct {
		name     string
		ipv4     bool
		ipv6     bool
		fallback bool
	}{
		{name: "IPv4", ipv4: true},
		{name: "IPv6", ipv6: true},
		{name: "IPv4 fallback", ipv4: true, fallback: true},
		{name: "IPv6 fallback", ipv6: true, fallback: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var output bytes.Buffer
			previousLogger := log.Logger
			log.Logger = zerolog.New(&output)
			t.Cleanup(func() {
				log.Logger = previousLogger
			})

			providers := []string{":first-provider", ":second-provider"}
			ddns := testApp("")
			family := "IPv4"
			if tc.ipv6 {
				family = "IPv6"
			}
			expectedLevel := "error"
			if tc.fallback {
				expectedLevel = "debug"
				ip := "203.0.113.42"
				network, address := "tcp4", "127.0.0.1:0"
				if tc.ipv6 {
					ip = "2001:db8::42"
					network, address = "tcp6", "[::1]:0"
				}
				listener, err := (&net.ListenConfig{}).Listen(context.Background(), network, address)
				if err != nil && tc.ipv6 {
					t.Skipf("IPv6 loopback unavailable: %v", err)
				}
				require.NoError(t, err)
				srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					_, _ = io.WriteString(w, ip)
				}))
				srv.Listener = listener
				srv.Start()
				t.Cleanup(srv.Close)
				providers = append(providers, srv.URL)

				// Prevent Route53 requests after a successful WAN lookup.
				ctx, cancel := context.WithCancelCause(context.Background())
				t.Cleanup(func() { cancel(nil) })
				ddns.r53, err = route53.New(ctx, "test", "test", "test", 0, 0)
				require.NoError(t, err)
				cancel(nil)
			}
			ddns.cfg.Route53.HandleIPv4 = new(tc.ipv4)
			ddns.cfg.Route53.HandleIPv6 = new(tc.ipv6)
			ddns.wip = wanip.New(
				wanip.WithIPv4Providers(providers),
				wanip.WithIPv6Providers(providers),
			)

			require.NotPanics(t, ddns.Run)

			decoder := json.NewDecoder(&output)
			var loggedProviders []string
			for {
				var record struct {
					Level       string `json:"level"`
					Message     string `json:"message"`
					Error       string `json:"error"`
					ProviderURL string `json:"provider-url"`
				}
				err := decoder.Decode(&record)
				if err == io.EOF {
					break
				}
				require.NoError(t, err)
				if record.ProviderURL == "" {
					continue
				}
				require.Equal(t, expectedLevel, record.Level)
				require.Equal(t, "Cannot retrieve WAN "+family+" address", record.Message)
				require.Contains(t, record.Error, record.ProviderURL)
				loggedProviders = append(loggedProviders, record.ProviderURL)
			}
			require.Equal(t, providers[:2], loggedProviders)
		})
	}
}

func TestStartWithoutScheduleReturns(t *testing.T) {
	t.Parallel()

	done := make(chan error, 1)
	go func() {
		done <- testApp("").Start(context.Background())
	}()

	select {
	case err := <-done:
		require.NoError(t, err)
	case <-time.After(time.Second):
		t.Fatal("Start did not return without a schedule")
	}
}

func TestStartWaitsForContextCancellation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	done := make(chan error, 1)
	go func() {
		done <- testApp("* * * * *").Start(ctx)
	}()

	select {
	case err := <-done:
		t.Fatalf("Start returned before shutdown: %v", err)
	case <-time.After(100 * time.Millisecond):
	}

	cancel(context.Canceled)

	select {
	case err := <-done:
		require.NoError(t, err)
	case <-time.After(time.Second):
		t.Fatal("Start did not stop after context cancellation")
	}
}

func testApp(schedule string) *DDNSRoute53 {
	return &DDNSRoute53{
		cfg: &config.Config{
			Cli: config.Cli{
				Schedule: schedule,
			},
			Route53: &config.Route53{
				HandleIPv4: new(false),
				HandleIPv6: new(false),
			},
		},
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor),
		)),
	}
}
