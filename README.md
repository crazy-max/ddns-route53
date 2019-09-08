<p align="center">
  <a href="https://github.com/crazy-max/ddns-route53/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/ddns-route53.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/ddns-route53/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/ddns-route53/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/badge/dynamic/json.svg?label=version&query=$.results[1].name&url=https://hub.docker.com/v2/repositories/crazymax/ddns-route53/tags&style=flat-square" alt="Latest Version"></a>
  <a href="https://github.com/crazy-max/ddns-route53/actions"><img src="https://github.com/crazy-max/ddns-route53/workflows/build/badge.svg" alt="Build Status"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/docker/stars/crazymax/ddns-route53.svg?style=flat-square" alt="Docker Stars"></a>
  <a href="https://hub.docker.com/r/crazymax/ddns-route53/"><img src="https://img.shields.io/docker/pulls/crazymax/ddns-route53.svg?style=flat-square" alt="Docker Pulls"></a>
  <br /><a href="https://quay.io/repository/crazymax/ddns-route53"><img src="https://quay.io/repository/crazymax/ddns-route53/status?style=flat-square" alt="Docker Repository on Quay"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/ddns-route53"><img src="https://goreportcard.com/badge/github.com/crazy-max/ddns-route53?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/ddns-route53"><img src="https://img.shields.io/codacy/grade/93db381dca8b441cb69b45b75f5e10ed.svg?style=flat-square" alt="Code Quality"></a>
  <br /><a href="https://www.patreon.com/crazymax"><img src="https://img.shields.io/badge/donate-patreon-f96854.svg?logo=patreon&style=flat-square" alt="Support me on Patreon"></a>
  <a href="https://www.paypal.me/crazyws"><img src="https://img.shields.io/badge/donate-paypal-00457c.svg?logo=paypal&style=flat-square" alt="Donate Paypal"></a>
</p>

## About

**ddns-route53** :motorway: is a CLI application written in [Go](https://golang.org/) that lets you run your own [dynamic DNS](https://en.wikipedia.org/wiki/Dynamic_DNS) service with [Amazon Route 53](https://aws.amazon.com/route53/) on a time-based schedule.

## Features

* Handle IPv4 and IPv6 addresses
* Internal cron implementation through go routines
* :whale: Official [Docker image available](doc/install/docker.md)

## Documentation

* [Prerequisites](doc/prerequisites.md)
* Install
  * [With Docker](doc/install/docker.md)
  * [From binary](doc/install/binary.md)
  * [Linux service](doc/install/linux-service.md)
* [Usage](doc/usage.md)
* [Configuration](doc/configuration.md)

## How can I help ?

All kinds of contributions are welcome :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
But we're not gonna lie to each other, I'd rather you buy me a beer or two :beers:!

[![Support me on Patreon](.res/patreon.png)](https://www.patreon.com/crazymax) 
[![Paypal Donate](.res/paypal.png)](https://www.paypal.me/crazyws)

## License

MIT. See `LICENSE` for more details.
