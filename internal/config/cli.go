package config

import (
	"time"

	"github.com/alecthomas/kong"
)

// Cli holds command line args, flags and cmds
type Cli struct {
	Version         kong.VersionFlag
	Cfgfile         string        `kong:"name='config',env='CONFIG',help='ddns-route53 configuration file.'"`
	Schedule        string        `kong:"name='schedule',env='SCHEDULE',help='CRON expression format.'"`
	Ifname          string        `kong:"name='ifname',env='IFNAME',help='Network interface name to be used for WAN IP retrieval. Leave empty to use the default one.'"`
	MaxRetries      int           `kong:"name='max-retries',env='MAX_RETRIES',default='3',help='Number of retries in case of WAN IP retrieval and AWS request failure.'"`
	MaxBackoffDelay time.Duration `kong:"name='max-backoff-delay',env='MAX_BACKOFF_DELAY',default='5s',help='Max back off delay that is allowed to occur between retrying a failed AWS request.'"`
	LogLevel        string        `kong:"name='log-level',env='LOG_LEVEL',default='info',help='Set log level.'"`
	LogJSON         bool          `kong:"name='log-json',env='LOG_JSON',default='false',help='Enable JSON logging output.'"`
	LogCaller       bool          `kong:"name='log-caller',env='LOG_CALLER',default='false',help='Add file:line of the caller to log output.'"`
	LogNoColor      bool          `kong:"name='log-nocolor',env='LOG_NOCOLOR',default='false',help='Disables the colorized output.'"`
}
