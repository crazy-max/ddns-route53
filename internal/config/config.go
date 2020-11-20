package config

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/crazy-max/ddns-route53/v2/pkg/utl"
	"github.com/crazy-max/gonfig"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Config holds configuration details
type Config struct {
	Cli         Cli          `yaml:"-" json:"-" label:"-" file:"-"`
	Credentials *Credentials `yaml:"credentials,omitempty" json:"credentials,omitempty" validate:"omitempty"`
	Route53     *Route53     `yaml:"route53,omitempty" json:"route53,omitempty" validate:"required"`
	Ipprovider  string       `yaml:"ipProvider,omitempty" json:"ipProvider,omitempty"`
}

// Load returns Configuration struct
func Load(cli Cli) (*Config, error) {
	cfg := Config{
		Cli: cli,
	}

	fileLoader := gonfig.NewFileLoader(gonfig.FileLoaderConfig{
		Filename: cli.Cfgfile,
		Finder: gonfig.Finder{
			BasePaths:  []string{"/etc/ddns-route53/ddns-route53", "$XDG_CONFIG_HOME/ddns-route53", "$HOME/.config/ddns-route53", "./ddns-route53"},
			Extensions: []string{"yaml", "yml"},
		},
	})
	if found, err := fileLoader.Load(&cfg); err != nil {
		return nil, errors.Wrap(err, "Failed to decode configuration from file")
	} else if !found {
		log.Debug().Msg("No configuration file found")
	} else {
		log.Info().Msgf("Configuration loaded from file: %s", fileLoader.GetFilename())
	}

	envLoader := gonfig.NewEnvLoader(gonfig.EnvLoaderConfig{
		Prefix: "DDNSR53_",
	})
	if found, err := envLoader.Load(&cfg); err != nil {
		return nil, errors.Wrap(err, "Failed to decode configuration from environment variables")
	} else if !found {
		log.Debug().Msg("No DDNSR53_* environment variables defined")
	} else {
		log.Info().Msgf("Configuration loaded from %d environment variables", len(envLoader.GetVars()))
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) validate() error {
	if cfg.Route53 == nil {
		return errors.New("route53 configuration required")
	}

	if cfg.Route53.RecordsSet == nil || len(cfg.Route53.RecordsSet) == 0 {
		return errors.New("empty record set")
	}

	for _, rs := range cfg.Route53.RecordsSet {
		if rs.Type == route53.RRTypeA {
			cfg.Route53.HandleIPv4 = utl.NewTrue()
		} else if rs.Type == route53.RRTypeAaaa {
			cfg.Route53.HandleIPv6 = utl.NewTrue()
		}
	}

	return validator.New().Struct(cfg)
}

// String returns the string representation of configuration
func (cfg *Config) String() string {
	b, _ := json.MarshalIndent(cfg, "", "  ")
	return string(b)
}
