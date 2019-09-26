# Changelog

## 1.4.0 (2019/09/26)

* Run container as non-root user

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

* Remove multi-platform Docker images, need rework (Issue #4)

## 1.3.0 (2019/08/31)

* Add multi-platform Docker images (Issue #4)
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
