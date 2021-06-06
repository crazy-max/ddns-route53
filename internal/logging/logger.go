package logging

import (
	"io"
	"os"
	"time"

	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure configures logger
func Configure(cli config.Cli) {
	var err error
	var w io.Writer

	// Adds support for NO_COLOR. More info https://no-color.org/
	_, noColor := os.LookupEnv("NO_COLOR")

	if !cli.LogJSON {
		w = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    noColor || cli.LogNoColor,
			TimeFormat: time.RFC1123,
		}
	} else {
		w = os.Stdout
	}

	ctx := zerolog.New(w).With().Timestamp()
	if cli.LogCaller {
		ctx = ctx.Caller()
	}

	log.Logger = ctx.Logger()

	logLevel, err := zerolog.ParseLevel(cli.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unknown log level")
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}
}
