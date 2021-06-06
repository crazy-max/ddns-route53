package config

import "github.com/alecthomas/kong"

// Cli holds command line args, flags and cmds
type Cli struct {
	Version    kong.VersionFlag
	Cfgfile    string `kong:"name='config',env='CONFIG',help='ddns-route53 configuration file.'"`
	Schedule   string `kong:"name='schedule',env='SCHEDULE',help='CRON expression format.'"`
	MaxRetries int    `kong:"name='max-retries',env='MAX_RETRIES',default='3',help='Number of retries in case of WAN IP retrieval failure.'"`
	LogLevel   string `kong:"name='log-level',env='LOG_LEVEL',default='info',help='Set log level.'"`
	LogJSON    bool   `kong:"name='log-json',env='LOG_JSON',default='false',help='Enable JSON logging output.'"`
	LogCaller  bool   `kong:"name='log-caller',env='LOG_CALLER',default='false',help='Add file:line of the caller to log output.'"`
	LogNoColor bool   `kong:"name='log-nocolor',env='LOG_NOCOLOR',default='false',help='Disables the colorized output.'"`
}
