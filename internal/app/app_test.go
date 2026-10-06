package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/crazy-max/ddns-route53/v2/pkg/wanip"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestRunLogsEachWANProviderFailure(t *testing.T) {
	for _, tc := range []struct {
		name string
		ipv4 bool
		ipv6 bool
	}{
		{name: "IPv4", ipv4: true},
		{name: "IPv6", ipv6: true},
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
			ddns.cfg.Route53.HandleIPv4 = new(tc.ipv4)
			ddns.cfg.Route53.HandleIPv6 = new(tc.ipv6)
			ddns.wip = wanip.New(
				wanip.WithIPv4Providers(providers),
				wanip.WithIPv6Providers(providers),
			)

			require.NotPanics(t, ddns.Run)

			decoder := json.NewDecoder(&output)
			for _, provider := range providers {
				var record struct {
					Level       string `json:"level"`
					Message     string `json:"message"`
					Error       string `json:"error"`
					ProviderURL string `json:"provider-url"`
				}
				require.NoError(t, decoder.Decode(&record))
				require.Equal(t, "error", record.Level)
				require.Equal(t, "Cannot retrieve WAN "+tc.name+" address", record.Message)
				require.Equal(t, provider, record.ProviderURL)
				require.Contains(t, record.Error, provider)
			}
			var extraRecord any
			require.ErrorIs(t, decoder.Decode(&extraRecord), io.EOF)
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
