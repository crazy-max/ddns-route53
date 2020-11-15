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
	runtime.GOMAXPROCS(runtime.NumCPU())

	meta.Version = version
	meta.UserAgent = fmt.Sprintf("%s/%s go/%s %s", meta.ID, meta.Version, runtime.Version()[2:], strings.Title(runtime.GOOS))

	// Parse command line
	_ = kong.Parse(&cli,
		kong.Name(meta.ID),
		kong.Description(fmt.Sprintf("%s. More info: %s", meta.Desc, meta.URL)),
		kong.UsageOnError(),
		kong.Vars{
			"version": fmt.Sprintf("%s", version),
		},
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	// Init
	logging.Configure(cli)
	log.Info().Str("version", version).Msgf("Starting %s", meta.Name)

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
	cfg, err := config.Load(cli)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load configuration")
	}
	log.Debug().Msg(cfg.String())

	// Init
	if ddnsRoute53, err = app.New(meta, cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize ddns-route53")
	}

	// Start
	if err = ddnsRoute53.Start(); err != nil {
		log.Fatal().Err(err).Msg("Cannot start ddns-route53")
	}
}
