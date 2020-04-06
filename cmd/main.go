package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
	"github.com/crazy-max/ddns-route53/internal/app"
	"github.com/crazy-max/ddns-route53/internal/config"
	"github.com/crazy-max/ddns-route53/internal/logging"
	"github.com/rs/zerolog/log"
)

var (
	ddnsRoute53 *app.Client
	cli         config.Cli
	version     = "dev"
)

func main() {
	// Parse command line
	_ = kong.Parse(&cli,
		kong.Name("ddns-route53"),
		kong.Description(`Dynamic DNS for Amazon Route 53‎ on a time-based schedule. More info: https://github.com/crazy-max/ddns-route53`),
		kong.UsageOnError(),
		kong.Vars{
			"version": fmt.Sprintf("%s", version),
		},
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	// Load timezone location
	location, err := time.LoadLocation(cli.Timezone)
	if err != nil {
		log.Panic().Err(err).Msgf("Cannot load timezone %s", cli.Timezone)
	}

	// Init
	logging.Configure(&cli, location)
	log.Info().Msgf("Starting ddns-route53 %s", version)

	// Handle os signals
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-channel
		ddnsRoute53.Close()
		log.Warn().Msgf("Caught signal %v", sig)
		os.Exit(0)
	}()

	// Load and check configuration
	cfg, err := config.Load(cli, version)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load configuration")
	}
	if err := cfg.Check(); err != nil {
		log.Fatal().Err(err).Msg("Improper configuration")
	}

	// Init
	if ddnsRoute53, err = app.New(cfg, location); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize ddns-route53")
	}

	// Start
	if err = ddnsRoute53.Start(); err != nil {
		log.Fatal().Err(err).Msg("Cannot start ddns-route53")
	}
}
