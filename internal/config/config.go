package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/service/route53"
	"gopkg.in/yaml.v2"
)

// Configuration holds configuration details
type Configuration struct {
	Flags       Flags
	App         App
	Credentials Credentials `yaml:"credentials,omitempty"`
	Route53     Route53     `yaml:"route53,omitempty"`
}

// Flags holds flags from command line
type Flags struct {
	Cfgfile  string
	Schedule string
	Timezone string
	LogLevel string
	LogJson  bool
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
func Load(fl Flags, version string) (*Configuration, error) {
	var err error
	var cfg = Configuration{
		Flags: fl,
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

	if _, err = os.Lstat(fl.Cfgfile); err != nil {
		return nil, fmt.Errorf("unable to open config file, %s", err)
	}

	bytes, err := ioutil.ReadFile(fl.Cfgfile)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file, %s", err)
	}

	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return &cfg, nil
}

// Check verifies Configuration values
func (cfg *Configuration) Check() error {
	if cfg.Credentials.AccessKeyID == "" {
		return errors.New("AWS Access ID is required")
	}

	if cfg.Credentials.SecretAccessKey == "" {
		return errors.New("AWS Secret Access Key is required")
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
		cfg.Route53.HandleIPv4 = rs.Type == route53.RRTypeA
		cfg.Route53.HandleIPv6 = rs.Type == route53.RRTypeAaaa
	}

	return nil
}
