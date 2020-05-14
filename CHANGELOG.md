# Changelog

## 1.9.0 (2020/05/14)

* Allow to use `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_HOSTED_ZONE_ID` environment variables (#130)
* Check AWS Route53 hosted zone ID entry
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.30.27

## 1.8.0 (2020/05/06)

* Add `--log-caller` flag
* Flag `--log-json` not handled
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.30.20
* Update deps

## 1.7.0 (2020/04/06)

* Switch to kong command-line parser
* Use Open Container Specification labels as label-schema.org ones are deprecated
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.30.3
* Update deps

## 1.6.0 (2019/12/20)

* Strengthen WAN IP address retrieval and validation
* Move ident.me client to pkg
* Add `--max-retries` flag
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.26.6
* Update deps
* Go 1.13.5

## 1.5.0 (2019/11/11)

* Seconds field is now optional
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.25.31
* Cache go modules

## 1.4.2 (2019/10/08)

* Allow both IPv4 and IPv6 records (#7)
* Use Docker buildx action v1

## 1.4.1 (2019/09/27)

* Build on respective platforms
* Update deps
* Go 1.12.10

## 1.4.0 (2019/09/26)

* Run Docker container as non-root user

## 1.3.4 (2019/09/22)

* Use GOPROXY
* Update workflow
  * Use softprops/action-gh-release to create GitHub release
  * Use [ghaction-goreleaser](https://github.com/crazy-max/ghaction-goreleaser) GitHub Action
  * Use [ghaction-docker-buildx](https://github.com/crazy-max/ghaction-docker-buildx) GitHub Action
* Stop publishing Docker image on Quay

## 1.3.3 (2019/09/08)

* Fix DockerHub latest tag

## 1.3.2 (2019/08/31)

* Reenable multi-platform Docker images (Issue #4)

## 1.3.1 (2019/08/31)

* Remove multi-platform Docker image, need rework (Issue #4)

## 1.3.0 (2019/08/31)

* Multi-platform Docker image (Issue #4)
* Remove GitHub Package support

## 1.2.1 (2019/08/29)

* Push Docker image to GitHub Package

## 1.2.0 (2019/08/29)

* Switch to GitHub Actions

## 1.1.0 (2019/07/18)

* Display next execution time
* Use v3 robfig/cron
* Go 1.12.4

## 1.0.0 (2019/06/03)

* Run once on startup (Issue #2)
* Update libs

## 0.2.0 (2019/04/02)

* Handle IPv4 or IPv6 only if included in a record set (Issue #1)
* Update libs

## 0.1.1 (2019/04/02)

* Fix update of last IPv4/IPv6

## 0.1.0 (2019/04/01)

* Initial version based on [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.19.6
