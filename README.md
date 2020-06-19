<p align="center">
  <a href="https://github.com/crazy-max/ddns-route53/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/ddns-route53.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/ddns-route53/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/ddns-route53/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://github.com/crazy-max/ddns-route53/actions?workflow=build"><img src="https://img.shields.io/github/workflow/status/crazy-max/ddns-route53/build?label=build&logo=github&style=flat-square" alt="Build Status"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/docker/stars/crazymax/ddns-route53.svg?style=flat-square&logo=docker" alt="Docker Stars"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/docker/pulls/crazymax/ddns-route53.svg?style=flat-square&logo=docker" alt="Docker Pulls"></a>
  <br /><a href="https://goreportcard.com/report/github.com/crazy-max/ddns-route53"><img src="https://goreportcard.com/badge/github.com/crazy-max/ddns-route53?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/ddns-route53"><img src="https://img.shields.io/codacy/grade/93db381dca8b441cb69b45b75f5e10ed.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://github.com/sponsors/crazy-max"><img src="https://img.shields.io/badge/sponsor-crazy--max-181717.svg?logo=github&style=flat-square" alt="Become a sponsor"></a>
  <a href="https://www.paypal.me/crazyws"><img src="https://img.shields.io/badge/donate-paypal-00457c.svg?logo=paypal&style=flat-square" alt="Donate Paypal"></a>
</p>

## About

**ddns-route53** :motorway: is a CLI application written in [Go](https://golang.org/) that lets you run your own [dynamic DNS](https://en.wikipedia.org/wiki/Dynamic_DNS) service with [Amazon Route 53](https://aws.amazon.com/route53/) on a time-based schedule.

## Features

* Handle IPv4 and IPv6 addresses
* Internal cron implementation through go routines
* Official [Docker image available](doc/install/docker.md)

## Documentation

* [Prerequisites](doc/prerequisites.md)
* Install
  * [With Docker](doc/install/docker.md)
  * [From binary](doc/install/binary.md)
  * [Linux service](doc/install/linux-service.md)
  * [On a UniFi Security Gateway](doc/install/unifi.md)
* [Usage](doc/usage.md)
* [Configuration](doc/configuration.md)

## How can I help?

All kinds of contributions are welcome :raised_hands:! The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon: You can also support this project by [**becoming a sponsor on GitHub**](https://github.com/sponsors/crazy-max) :clap: or by making a [Paypal donation](https://www.paypal.me/crazyws) to ensure this journey continues indefinitely! :rocket:

Thanks again for your support, it is much appreciated! :pray:

## License

MIT. See `LICENSE` for more details.
