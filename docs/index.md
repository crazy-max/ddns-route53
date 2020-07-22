<img src="assets/logo.png" alt="ddns-route53" width="128px" style="display: block; margin-left: auto; margin-right: auto"/>

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

---

## What is ddns-route53?

**ddns-route53** :motorway: is a CLI application written in [Go](https://golang.org/) and delivered as a
[single executable]({{ config.repo_url }}releases/latest) (and a
[Docker image](https://hub.docker.com/r/crazymax/ddns-route53/)) that lets you run your own
[dynamic DNS](https://en.wikipedia.org/wiki/Dynamic_DNS) service with [Amazon Route 53](https://aws.amazon.com/route53/)
on a time-based schedule.

With Go, this can be done with an independent binary distribution across all platforms and architectures that Go supports.
This support includes Linux, macOS, and Windows, on architectures like amd64, i386, ARM, PowerPC, and others.

## Features

* Handle IPv4 and IPv6 addresses
* Internal cron implementation through go routines

## License

This project is licensed under the terms of the MIT license.
