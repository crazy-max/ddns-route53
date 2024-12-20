# Command Line

## Usage

```shell
ddns-route53 [options]
```

## Options

```
$ ddns-route53 --help
Usage: ddns-route53

Dynamic DNS for Amazon Route 53 on a time-based schedule. More info:
https://github.com/crazy-max/ddns-route53

Flags:
  -h, --help                    Show context-sensitive help.
      --version
      --config=STRING           ddns-route53 configuration file ($CONFIG).
      --schedule=STRING         CRON expression format ($SCHEDULE).
      --ifname=STRING           Network interface name to be used for WAN IP
                                retrieval. Leave empty to use the default one
                                ($IFNAME).
      --max-retries=3           Number of retries in case of WAN IP retrieval
                                and AWS request failure ($MAX_RETRIES).
      --max-backoff-delay=5s    Max back off delay that is allowed to
                                occur between retrying a failed AWS request
                                ($MAX_BACKOFF_DELAY).
      --log-level="info"        Set log level ($LOG_LEVEL).
      --log-json                Enable JSON logging output ($LOG_JSON).
      --log-caller              Add file:line of the caller to log output
                                ($LOG_CALLER).
      --log-nocolor             Disables the colorized output ($LOG_NOCOLOR).
```

## Environment variables

The following environment variables can be used in place:

| Name                | Default | Description                                                                                                     |
|---------------------|---------|-----------------------------------------------------------------------------------------------------------------|
| `CONFIG`            |         | ddns-route53 configuration file                                                                                 |
| `SCHEDULE`          |         | [CRON expression](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format) to schedule ddns-route53 |
| `IFNAME`            |         | Network interface name to be used for WAN IP retrieval. Leave empty to use the default one                      |
| `MAX_RETRIES`       | `3`     | Number of retries in case of WAN IP retrieval and AWS request failure                                           |
| `MAX_BACKOFF_DELAY` | `5s`    | Max back off delay that is allowed to occur between retrying a failed AWS request                               |
| `LOG_LEVEL`         | `info`  | Log level output                                                                                                |
| `LOG_JSON`          | `false` | Enable JSON logging output                                                                                      |
| `LOG_CALLER`        | `false` | Enable to add `file:line` of the caller                                                                         |
| `LOG_NOCOLOR`       | `false` | Disables the colorized output                                                                                   |

!!! warning
    If schedule is not set, the app will run once and then exit.
