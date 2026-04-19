package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	_ "time/tzdata"

	"github.com/alecthomas/kong"
	"github.com/crazy-max/ddns-route53/v2/internal/app"
	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/crazy-max/ddns-route53/v2/internal/logging"
	"github.com/crazy-max/ddns-route53/v2/internal/model"
	"github.com/rs/zerolog/log"
)

var (
	ddnsRoute53 *app.DDNSRoute53
	cli         config.Cli
	version     = "dev"
	meta        = model.Meta{
		ID:     "ddns-route53",
		Name:   "ddns-route53",
		Desc:   "Dynamic DNS for Amazon Route 53 on a time-based schedule",
		URL:    "https://github.com/crazy-max/ddns-route53",
		Logo:   "https://raw.githubusercontent.com/crazy-max/ddns-route53/master/.github/ddns-route53.png",
		Author: "CrazyMax",
	}
)

func main() {
	var err error

	meta.Version = version
	meta.UserAgent = fmt.Sprintf("%s/%s go/%s %s", meta.ID, meta.Version, runtime.Version()[2:], strings.Title(runtime.GOOS)) //nolint:staticcheck // ignoring "SA1019: strings.Title is deprecated", as for our use we don't need full unicode support

	_ = kong.Parse(&cli,
		kong.Name(meta.ID),
		kong.Description(fmt.Sprintf("%s. More info: %s", meta.Desc, meta.URL)),
		kong.UsageOnError(),
		kong.Vars{
			"version": version,
		},
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	logging.Configure(cli)
	log.Info().Str("version", version).Msgf("Starting %s", meta.Name)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-channel
		ddnsRoute53.Close()
		log.Warn().Msgf("Caught signal %v", sig)
		os.Exit(0)
	}()

	cfg, err := config.Load(cli)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load configuration")
	}
	log.Debug().Interface("config", cfg).Msg("Configuration")

	if ddnsRoute53, err = app.New(meta, cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize ddns-route53")
	}

	if err = ddnsRoute53.Start(); err != nil {
		log.Fatal().Err(err).Msg("Cannot start ddns-route53")
	}
}
