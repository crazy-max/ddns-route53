<p align="center">
  <a href="https://github.com/crazy-max/ddns-route53/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/ddns-route53.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/ddns-route53/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/ddns-route53/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/badge/dynamic/json.svg?label=version&query=$.results[1].name&url=https://hub.docker.com/v2/repositories/crazymax/ddns-route53/tags&style=flat-square" alt="Latest Version"></a>
  <a href="https://travis-ci.com/crazy-max/ddns-route53"><img src="https://img.shields.io/travis/com/crazy-max/ddns-route53/master.svg?style=flat-square" alt="Build Status"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/docker/stars/crazymax/ddns-route53.svg?style=flat-square" alt="Docker Stars"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/docker/pulls/crazymax/ddns-route53.svg?style=flat-square" alt="Docker Pulls"></a>
  <br /><a href="https://quay.io/repository/crazymax/ddns-route53"><img src="https://quay.io/repository/crazymax/ddns-route53/status?style=flat-square" alt="Docker Repository on Quay"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/ddns-route53"><img src="https://goreportcard.com/badge/github.com/crazy-max/ddns-route53?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/ddns-route53"><img src="https://img.shields.io/codacy/grade/93db381dca8b441cb69b45b75f5e10ed.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=KLCPAAPLKWJAG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

**ddns-route53** :motorway: is a CLI application written in [Go](https://golang.org/) that lets you run your own [dynamic DNS](https://en.wikipedia.org/wiki/Dynamic_DNS) service with [Amazon Route 53](https://aws.amazon.com/route53/) on a time-based schedule.

## Features

* Handle IPv4 and IPv6 addresses
* Internal cron implementation through go routines
* :whale: Official [Docker image available](#docker)

## Prerequisites

I assume you have already created a [Route 53 Hosted Zones](https://console.aws.amazon.com/route53/home#hosted-zones:) as a **Public Hosted Zone** type and setted Amazon name servers in your domain name registrar.

Go to the [IAM Policies page](https://console.aws.amazon.com/iam/home#/policies) and click on **Create Policy**.

Then click on **JSON** tab and paste the following content :

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "route53:ChangeResourceRecordSets"
            ],
            "Effect": "Allow",
            "Resource": "arn:aws:route53:::hostedzone/<HOSTED_ZONE_ID>"
        }
    ]
}
```

> Replace `<HOSTED_ZONE_ID>` with your current hosted zone id you can find on [Route 53 Hosted Zones page](https://console.aws.amazon.com/route53/home#hosted-zones:).

Enter a **Policy Name** and click **Create Policy**.

Now go to the [IAM Users page](https://console.aws.amazon.com/iam/home#/users) and click the **Add user** button.

Enter a **User name**, check **Programmatic access** for _Access type_ and click **Next: Permissions**.

Choose the last option **Attach existing policies directly** and fill in the **Search** field with the name of the policy you created before and click **Next: Review** then **Create user**.

An _Access Key ID_ and a _Secret Access key_ will be displayed. This is the credentials needed for ddns-route53. Save them somewhere since you will need them in the [configuration](#configuration) step.

## Download

ddns-route53 binaries are available in [releases](https://github.com/crazy-max/ddns-route53/releases) page.

Choose the archive matching the destination platform and extract ddns-route53:

```
$ cd /opt
$ wget -qO- https://github.com/crazy-max/ddns-route53/releases/download/v1.0.0/ddns-route53_1.0.0_linux_x86_64.tar.gz | tar -zxvf - ddns-route53
```

After getting the binary, it can be tested with `./ddns-route53 --help` or moved to a permanent location.

```
$ ./ddns-route53 --help
usage: ddns-route53 --config=CONFIG [<flags>]

Dynamic DNS for Amazon Route 53‎ on a time-based schedule. More info on
https://github.com/crazy-max/ddns-route53

Flags:
  --help               Show context-sensitive help (also try --help-long and
                       --help-man).
  --config=CONFIG      ddns-route53 configuration file.
  --schedule=SCHEDULE  CRON expression format.
  --timezone="UTC"     Timezone assigned to ddns-route53.
  --log-level="info"   Set log level.
  --log-json           Enable JSON logging output.
  --version            Show application version.
```

## Usage

`ddns-route53 --config=CONFIG [<flags>]`

* `--help` : Show help text and exit. _Optional_.
* `--version` : Show version and exit. _Optional_.
* `--config <path>` : ddns-route53 YAML configuration file. **Required**. (example: `ddns-route53.yml`).
* `--schedule <cron expression>` : [CRON expression](https://godoc.org/github.com/crazy-max/cron#hdr-CRON_Expression_Format) to schedule ddns-route53. _Optional_. (example: `0 */30 * * * *`).
* `--timezone <timezone>` : Timezone assigned to ddns-route53. _Optional_. (default: `UTC`).
* `--log-level <level>` : Log level output. _Optional_. (default: `info`).
* `--log-json` : Enable JSON logging output. _Optional_. (default: `false`).

## Configuration

Before running ddns-route53, you must create your first configuration file. Here is a YAML structure example :

```yml
credentials:
  access_key_id: "<AWS_ACCESS_KEY_ID>"
  secret_access_key: "<AWS_SECRET_ACCESS_KEY>"

route53:
  hosted_zone_id: "<HOSTED_ZONE_ID>"
  records_set:
    - name: "ddns.example.com."
      type: "A"
      ttl: 300
    - name: "another.example2.com."
      type: "A"
      ttl: 300
```

> Replace `<AWS_ACCESS_KEY_ID>`, `<AWS_SECRET_ACCESS_KEY>` and `<HOSTED_ZONE_ID>` with the correponded values.

* `credentials`
  * `access_key_id`: AWS Access Key.
  * `secret_access_key`: AWS Secret Key.
* `route53`
  * `hosted_zone_id`: AWS Route53 hosted zone ID.
  * `records_set`: Map of records set.
    * `name`: AWS Route53 record set name (don't forget to add a dot at the end).
    * `type`: AWS Route53 record set type, can be `A` or `AAAA`.
    * `ttl`: AWS Route53 record TTL (time to live) in seconds.

## Docker

ddns-route53 provides automatically updated Docker :whale: images within [Docker Hub](https://hub.docker.com/r/crazymax/ddns-route53) and [Quay](https://quay.io/repository/crazymax/ddns-route53). It is possible to always use the latest stable tag or to use another service that handles updating Docker images.

Environment variables can be used within your container :

* `TZ` : Timezone assigned to ddns-route53
* `SCHEDULE` : [CRON expression](https://godoc.org/github.com/crazy-max/cron#hdr-CRON_Expression_Format) to schedule ddns-route53
* `LOG_LEVEL` : Log level output (default `info`)
* `LOG_JSON`: Enable JSON logging output (default `false`)

Docker compose is the recommended way to run this image. Copy the content of folder [.res/compose](.res/compose) in `/opt/ddns-route53/` on your host for example. Edit the compose and config file with your preferences and run the following commands :

```bash
docker-compose up -d
docker-compose logs -f
```

Or use the following command :

```bash
$ docker run -d --name ddns-route53 \
  -e "TZ=Europe/Paris" \
  -e "SCHEDULE=0 */30 * * * *" \
  -e "LOG_LEVEL=info" \
  -e "LOG_JSON=false" \
  -v "$(pwd)/ddns-route53.yml:/ddns-route53.yml:ro" \
  crazymax/ddns-route53:latest
```

## How can I help ?

All kinds of contributions are welcome :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
But we're not gonna lie to each other, I'd rather you buy me a beer or two :beers:!

[![Paypal](.res/paypal-donate.png)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=KLCPAAPLKWJAG)

## License

MIT. See `LICENSE` for more details.
