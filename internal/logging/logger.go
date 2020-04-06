package logging

import (
	"os"
	"time"

	"github.com/crazy-max/ddns-route53/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure configures logger
func Configure(cli *config.Cli, location *time.Location) {
	var err error

	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(location)
	}

	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC1123,
	}).With().Timestamp().Logger()

	logLevel, err := zerolog.ParseLevel(cli.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unknown log level")
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}
}
