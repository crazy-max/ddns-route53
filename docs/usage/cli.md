# Command Line

## Usage

```shell
$ ddns-route53 [options]
```

## Options

```
$ ddns-route53 --help
Usage: ddns-route53

Dynamic DNS for Amazon Route 53 on a time-based schedule. More info:
https://github.com/crazy-max/ddns-route53

Flags:
  -h, --help                Show context-sensitive help.
      --version
      --config=STRING       ddns-route53 configuration file ($CONFIG).
      --schedule=STRING     CRON expression format ($SCHEDULE).
      --max-retries=3       Number of retries in case of WAN IP retrieval
                            failure ($MAX_RETRIES).
      --log-level="info"    Set log level ($LOG_LEVEL).
      --log-json            Enable JSON logging output ($LOG_JSON).
      --log-caller          Add file:line of the caller to log output
                            ($LOG_CALLER).
```

## Environment variables

Following environment variables can be used in place:

| Name               | Default       | Description   |
|--------------------|---------------|---------------|
| `CONFIG`           |               | ddns-route53 configuration file |
| `SCHEDULE`         |               | [CRON expression](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format) to schedule ddns-route53 |
| `LOG_LEVEL`        | `info`        | Log level output |
| `LOG_JSON`         | `false`       | Enable JSON logging output |
| `LOG_CALLER`       | `false`       | Enable to add `file:line` of the caller |
