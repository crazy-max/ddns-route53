package main

import (
	"context"
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
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var version = "dev"

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func run() error {
	cli := config.Cli{}
	meta := model.Meta{
		ID:      "ddns-route53",
		Name:    "ddns-route53",
		Desc:    "Dynamic DNS for Amazon Route 53 on a time-based schedule",
		URL:     "https://github.com/crazy-max/ddns-route53",
		Logo:    "https://raw.githubusercontent.com/crazy-max/ddns-route53/master/.github/ddns-route53.png",
		Author:  "CrazyMax",
		Version: version,
	}

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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load(cli)
	if err != nil {
		return errors.Wrap(err, "cannot load configuration")
	}
	log.Debug().Interface("config", cfg).Msg("Configuration")

	ddnsRoute53, err := app.New(ctx, meta, cfg)
	if err != nil {
		return errors.Wrap(err, "cannot initialize ddns-route53")
	}

	if err := ddnsRoute53.Start(ctx); err != nil {
		return errors.Wrap(err, "cannot start ddns-route53")
	}

	if cause := context.Cause(ctx); cause != nil {
		log.Warn().Msg(strings.Title(cause.Error())) //nolint:staticcheck // ignoring "SA1019: strings.Title is deprecated", as for our use we don't need full unicode support
	}

	return nil
}
