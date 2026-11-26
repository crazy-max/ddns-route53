package config

import (
	"net/url"

	"github.com/crazy-max/gonfig"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// HostedZone holds AWS Route53 HostedZone data
type HostedZone struct {
	HostedZoneID string     `yaml:"hostedZoneID,omitempty" json:"hostedZoneID,omitempty" validate:"required"`
	RecordsSet   RecordsSet `yaml:"recordsSet,omitempty" json:"recordsSet,omitempty" validate:"dive"`
}

// Config holds configuration details
type Config struct {
	Cli Cli `yaml:"-" json:"-" label:"-" file:"-"`

	Credentials *Credentials `yaml:"credentials,omitempty" json:"credentials,omitempty" validate:"omitempty"`
	HostedZones []HostedZone `yaml:"hostedZones,omitempty" json:"hostedZones,omitempty" validate:"dive"`
	WanIP       *WanIP       `yaml:"wanip,omitempty" json:"wanip,omitempty" validate:"omitempty"`
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
	if len(cfg.HostedZones) == 0 {
		return errors.New("empty hosted zone list")
	}

	for _, hz := range cfg.HostedZones {
		if len(hz.RecordsSet) == 0 {
			return errors.New("empty record set")
		}
	}

	if cfg.WanIP != nil && cfg.WanIP.Providers != nil {
		if err := validateProviderURLs("ipv4", cfg.WanIP.Providers.IPv4); err != nil {
			return err
		}
		if err := validateProviderURLs("ipv6", cfg.WanIP.Providers.IPv6); err != nil {
			return err
		}
	}

	return validator.New().Struct(cfg)
}

func validateProviderURLs(family string, values []string) error {
	for i, raw := range values {
		u, err := url.Parse(raw)
		if err != nil {
			return errors.Wrapf(err, "invalid wanip.providers.%s[%d]", family, i)
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return errors.Errorf("invalid wanip.providers.%s[%d]: unsupported scheme %q", family, i, u.Scheme)
		}
		if u.Host == "" {
			return errors.Errorf("invalid wanip.providers.%s[%d]: host required", family, i)
		}
	}
	return nil
}
