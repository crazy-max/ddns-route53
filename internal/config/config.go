package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Configuration holds configuration details
type Configuration struct {
	Cli         Cli
	App         App
	Credentials Credentials `yaml:"credentials,omitempty"`
	Route53     Route53     `yaml:"route53,omitempty"`
}

// Cli holds command line args, flags and cmds
type Cli struct {
	Version    kong.VersionFlag
	Cfgfile    string `kong:"required,name='config',env='CONFIG',help='ddns-route53 configuration file.'"`
	Schedule   string `kong:"name='schedule',env='SCHEDULE',help='CRON expression format.'"`
	MaxRetries int    `kong:"name='max-retries',env='MAX_RETRIES',default='3',help='Number of retries in case of WAN IP retrieval failure.'"`
	Timezone   string `kong:"name='timezone',env='TZ',default='UTC',help='Timezone assigned to ddns-route53.'"`
	LogLevel   string `kong:"name='log-level',env='LOG_LEVEL',default='info',help='Set log level.'"`
	LogJSON    bool   `kong:"name='log-json',env='LOG_JSON',default='false',help='Enable JSON logging output.'"`
	LogCaller  bool   `kong:"name='log-caller',env='LOG_CALLER',default='false',help='Add file:line of the caller to log output.'"`
}

// App holds application details
type App struct {
	Name    string
	Desc    string
	URL     string
	Author  string
	Version string
}

// Credentials holds data necessary for AWS configuration
type Credentials struct {
	AccessKeyID     string `yaml:"access_key_id,omitempty"`
	SecretAccessKey string `yaml:"secret_access_key,omitempty"`
}

// Route53 holds AWS Route53 data
type Route53 struct {
	HostedZoneID string      `yaml:"hosted_zone_id,omitempty"`
	RecordsSet   []RecordSet `yaml:"records_set,omitempty"`
	HandleIPv4   bool
	HandleIPv6   bool
}

// RecordSet holds data necessary for record set configuration
type RecordSet struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type,omitempty"`
	TTL  int64  `yaml:"ttl,omitempty"`
}

// Load returns Configuration struct
func Load(cli Cli, version string) (*Configuration, error) {
	var err error
	var cfg = Configuration{
		Cli: cli,
		App: App{
			Name:    "ddns-route53",
			Desc:    "Dynamic DNS for Amazon Route 53‎ on a time-based schedule",
			URL:     "https://github.com/crazy-max/ddns-route53",
			Author:  "CrazyMax",
			Version: version,
		},
		Route53: Route53{
			HandleIPv4: false,
			HandleIPv6: false,
		},
	}

	if _, err = os.Lstat(cli.Cfgfile); err != nil {
		return nil, errors.Wrap(err, "unable to open config file")
	}

	bytes, err := ioutil.ReadFile(cli.Cfgfile)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config file")
	}

	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, errors.Wrap(err, "unable to decode into struct")
	}

	return &cfg, nil
}

// Check verifies Configuration values
func (cfg *Configuration) Check() error {
	cfg.Credentials.AccessKeyID = getEnv("AWS_ACCESS_KEY_ID", cfg.Credentials.AccessKeyID)
	if cfg.Credentials.AccessKeyID == "" {
		return errors.New("AWS Access ID is required")
	}

	cfg.Credentials.SecretAccessKey = getEnv("AWS_SECRET_ACCESS_KEY", cfg.Credentials.SecretAccessKey)
	if cfg.Credentials.SecretAccessKey == "" {
		return errors.New("AWS Secret Access Key is required")
	}

	cfg.Route53.HostedZoneID = getEnv("AWS_HOSTED_ZONE_ID", cfg.Route53.HostedZoneID)
	if cfg.Route53.HostedZoneID == "" {
		return errors.New("AWS Route53 hosted zone ID")
	}

	if len(cfg.Route53.RecordsSet) == 0 {
		return errors.New("empty record set")
	}

	for i, rs := range cfg.Route53.RecordsSet {
		if rs.Name == "" {
			return fmt.Errorf("missing record set name at index %d", i)
		}
		if rs.Type == "" {
			return fmt.Errorf("missing record set type for %s", rs.Name)
		}
		if rs.Type != route53.RRTypeA && rs.Type != route53.RRTypeAaaa {
			return fmt.Errorf("invalid record set type %s for %s", rs.Type, rs.Name)
		}
		if rs.TTL < 1 {
			return fmt.Errorf("invalid record set TTL %d for %s", rs.TTL, rs.Name)
		}
		if rs.Type == route53.RRTypeA {
			cfg.Route53.HandleIPv4 = true
		} else if rs.Type == route53.RRTypeAaaa {
			cfg.Route53.HandleIPv6 = true
		}
	}

	return nil
}

// getEnv retrieves the value of the environment variable named by the key
// or fallback if not found
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
