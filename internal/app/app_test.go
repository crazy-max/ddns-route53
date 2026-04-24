package app

import (
	"context"
	"testing"
	"time"

	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/require"
)

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
