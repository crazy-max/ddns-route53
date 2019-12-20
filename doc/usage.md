# Usage

## Command line

`ddns-route53 --config=CONFIG [<flags>]`

* `--help` : Show help text and exit. _Optional_.
* `--version` : Show version and exit. _Optional_.
* `--config <path>` : ddns-route53 YAML configuration file. **Required**. (example: `ddns-route53.yml`).
* `--schedule <cron expression>` : [CRON expression](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format) to schedule ddns-route53. _Optional_. (example: `*/30 * * * *`).
* `--max-retries <nb of retries>` : Number of retries in case of WAN IP retrieval failure. _Optional_. (example: `3`).
* `--timezone <timezone>` : Timezone assigned to ddns-route53. _Optional_. (default: `UTC`).
* `--log-level <level>` : Log level output. _Optional_. (default: `info`).
* `--log-json` : Enable JSON logging output. _Optional_. (default: `false`).
