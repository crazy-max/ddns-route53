# Changelog

## 2.13.0 (2024/12/20)

* Allow non-static AWS Credentials (#1204)
* Go 1.23 (#1258)
* Alpine Linux 3.21 (#1258)
* Switch to github.com/dromara/carbon/v2 2.5.2 (#1259)
* Bump github.com/alecthomas/kong to 1.6.0 (#1251)
* Bump github.com/aws/aws-sdk-go-v2 dependencies (#1254)
    * github.com/aws/aws-sdk-go-v2 to 1.32.7
    * github.com/aws/aws-sdk-go-v2/config to 1.28.7
    * github.com/aws/aws-sdk-go-v2/credentials to 1.17.48
    * github.com/aws/aws-sdk-go-v2/service/route53 to 1.46.4
* Bump github.com/go-playground/validator/v10 to 10.23.0 (#1219 #1253)
* Bump github.com/hashicorp/go-retryablehttp to 0.7.7 (#1214)
* Bump github.com/rs/zerolog to 1.33.0 (#1212)
* Bump golang.org/x/crypto to 0.31.0 (#1257)
* Bump golang.org/x/net to 0.23.0 (#1201)
* Bump golang.org/x/sys to 0.28.0 (#1217 #1252)

## 2.12.0 (2024/03/24)

* Generate sbom and provenance (#1193)
* Bump github.com/alecthomas/kong to 0.9.0 (#1187)
* Bump github.com/aws/aws-sdk-go-v2 dependencies (#1154 #1158 #1191)
    * github.com/aws/aws-sdk-go-v2 to 1.26.0
    * github.com/aws/aws-sdk-go-v2/config to 1.27.9
    * github.com/aws/aws-sdk-go-v2/credentials to 1.17.9
    * github.com/aws/aws-sdk-go-v2/service/route53 to 1.40.3
* Bump github.com/go-playground/validator/v10 to 10.19.0 (#1183)
* Bump github.com/rs/zerolog to 1.32.0 (#1172)
* Bump golang.org/x/crypto to 0.17.0 (#1153)
* Bump golang.org/x/sys to 0.18.0 (#1184)

## 2.11.0 (2023/12/16)

* Go 1.21 (#1152)
* Bump github.com/alecthomas/kong to 0.8.1 (#1077)
* Bump github.com/aws/aws-sdk-go-v2 to 1.24.0 (#1144)
* Bump github.com/aws/aws-sdk-go-v2/config to 1.26.1 (#1147)
* Bump github.com/aws/aws-sdk-go-v2/credentials to 1.16.12 (#1145)
* Bump github.com/aws/aws-sdk-go-v2/service/route53 to 1.35.5 (#1146)
* Bump github.com/go-playground/validator/v10 to 10.16.0 (#1096)
* Bump github.com/hashicorp/go-retryablehttp to 0.7.5 (#1101)
* Bump golang.org/x/net to 0.17.0 (#1076)
* Bump golang.org/x/sys to 0.15.0 (#1124)

## 2.10.0 (2023/09/27)

* Align AWS request retries with `max-retries` flag and decrease backoff delay to 5s (#1064)
* Add max backoff delay option (#1065)
* Go 1.20 (#995)
* Alpine Linux 3.18 (#996)
* Bump github.com/alecthomas/kong to 0.8.0 (#994)
* Bump github.com/aws/aws-sdk-go-v2 to 1.21.0 (#989 #1035)
* Bump github.com/aws/aws-sdk-go-v2/config to 1.18.42 (#908 #991 #1047 #1062)
* Bump github.com/aws/aws-sdk-go-v2/credentials to 1.13.37 (#992 #1046)
* Bump github.com/aws/aws-sdk-go-v2/service/route53 to 1.29.5 (#907 #993 #1037)
* Bump github.com/crazy-max/gonfig to 0.7.1 (#975)
* Bump github.com/go-playground/validator/v10 to 10.15.4 (#984 #1020 #1056)
* Bump github.com/hashicorp/go-retryablehttp to 0.7.4 (#905 #985)
* Bump github.com/rs/zerolog to 1.31.0 (#918 #961 #1011 #1063)
* Bump github.com/stretchr/testify to 1.8.4 (#937 #982)
* Bump golang.org/x/crypto to 0.1.0 (#939) 
* Bump golang.org/x/sys to 0.12.0 (#906 #938 #953 #969 #986 #997 #1019 #1044)
* Bump golang.org/x/text to 0.3.8 (#936)

## 2.9.0 (2022/12/31)

* Switch to retryable http client implementation (#900)
* Allow specifying custom interface (#901)
* Check record value with WAN IP and update accordingly (#902)
    * :warning: requires `route53:ListResourceRecordSets` authorization in IAM policy
* Update to AWS SDK v2 (#899)
* Go 1.19 (#846 #750)
* Alpine Linux 3.17 (#898)
* Enhance workflow (#847)
* Bump github.com/go-playground/validator/v10 from 10.10.0 to 10.11.1 (#667 #710 #825)
* Bump github.com/alecthomas/kong from 0.3.0 to 0.7.1 (#673 #749 #754 #864)
* Bump github.com/stretchr/testify from 1.7.0 to 1.8.1 (#675 #743 #767 #848)
* Bump github.com/rs/zerolog from 1.26.1 to 1.27.0 (#744)
* Bump golang.org/x/sys to v0.3.0 (#751 #903)
* Bump github.com/crazy-max/gonfig from 0.5.0 to 0.6.0 (#780)
* Bump github.com/rs/zerolog from 1.27.0 to 1.28.0 (#810)

## 2.8.0 (2022/01/26)

* Alpine Linux 3.15 (#622)
* MkDocs Material 8.1.8 (#637)
* goreleaser-xx 1.2.5 (#627)
* Move `syscall` to `golang.org/x/sys` (#612)
* Move from `io/ioutil` to `io` and `os` packages (#613)
* Enhance dockerfiles (#610 #611)
* Bump github.com/go-playground/validator/v10 from 10.9.0 to 10.10.0 (#614)
* Bump github.com/alecthomas/kong from 0.2.17 to 0.3.0 (#576 #599 #605 #620)
* Bump github.com/aws/aws-sdk-go from 1.40.37 to 1.42.41 (#551 #577 #615 #623 #636)
* Bump github.com/rs/zerolog from 1.24.0 to 1.26.1 (#534 #574 #607)

## 2.7.0 (2021/09/05)

* Go 1.17 (#519)
* Add `windows/arm64` artifact
* Wrong remaining time displayed
* Bump github.com/aws/aws-sdk-go from 1.38.69 to 1.40.37 (#518 #531)
* Bump github.com/rs/zerolog from 1.23.0 to 1.24.0 (#525)
* Bump github.com/crazy-max/gonfig from 0.4.0 to 0.5.0 (#520)
* Bump codecov/codecov-action from 1 to 2
* Bump github.com/go-playground/validator/v10 from 10.6.1 to 10.9.0 (#483 #508)
* Corrected Unifi documentation (#477)

## 2.6.1 (2021/06/29)

* `windows/arm64` not yet available

## 2.6.0 (2021/06/29)

* Bump github.com/aws/aws-sdk-go from 1.38.55 to 1.38.69 (#472)
* Add `linux/riscv64` and `windows/arm64` artifacts (#475)
* Alpine Linux 3.14 (#474)
* Added quotations to linux service (#473)
* Bump github.com/rs/zerolog from 1.22.0 to 1.23.0 (#463)
* Bump github.com/alecthomas/kong from 0.2.16 to 0.2.17 (#457)

## 2.5.0 (2021/06/05)

* Bump github.com/go-playground/validator/v10 from 10.4.1 to 10.6.1 (#435)
* Allow disabling log color output (#456)
* Update build workflow (#455)
* Add cni installation to unifi-os install (#438)
* MkDocs Materials 7.1.5 (#453)
* Fix artifacts download links
* Set `cacheonly` output for validators
* Move to `docker/metadata-action`
* Add `darwin/arm64` artifact
* Deploy docs on workflow dispatch or tag
* Bump github.com/rs/zerolog from 1.20.0 to 1.22.0 (#388 #437)
* Go 1.16 (#390)
* Bump github.com/aws/aws-sdk-go from 1.37.25 to 1.38.55 (#377 #391 #454)
* Bump github.com/alecthomas/kong from 0.2.15 to 0.2.16 (#376)

## 2.4.0 (2021/03/07)

* Bump github.com/aws/aws-sdk-go from 1.36.20 to 1.37.25 (#374)
* Fix WanIP provider
* Missing non-root user for Docker image
* Switch to `goreleaser-xx` (#373)
* Remove s390x Docker image support (#361)
* Bump github.com/alecthomas/kong from 0.2.12 to 0.2.15 (#353)
* Bump github.com/stretchr/testify from 1.6.1 to 1.7.0 (#339)

## 2.3.0 (2021/01/05)

* Bump github.com/aws/aws-sdk-go from 1.35.28 to 1.36.20 (#332)
* Refactor CI and dev workflow with buildx bake (#333)
      * Upload artifacts
      * Add `image-local` target
      * Single job for artifacts and image
      * Add `armv5`, `ppc64le` and `s390x` artifacts
* Handle multi IP providers (#331)
* Bump github.com/alecthomas/kong from 0.2.11 to 0.2.12 (#307)

## 2.2.0 (2020/11/15)

* Use embedded tzdata package
* Remove `--timezone` flag
* Docker image also available on [GitHub Container Registry](https://github.com/users/crazy-max/packages/container/package/ddns-route53)
* Switch to Docker actions
* Update deps

## 2.1.0 (2020/09/07)

* Don't fill record change if IP address not available (#224)
* Go 1.15
* Update deps

## 2.0.1 (2020/08/05)

* Fix nil pointer with AWS credentials (#186)

## 2.0.0 (2020/07/22)

:warning: See **Migration notes** in the documentation for breaking changes.

* Configuration transposed into environment variables
* `AWS_HOSTED_ZONE_ID` env var renamed `DDNSR53_ROUTE53_HOSTEDZONEID`
* Improve configuration validation
* All fields in configuration now _camelCased_
* Add tests and coverage
* Dockerfile enhanced
* Seek configuration file from default places
* Configuration file not required anymore
* Switch to [gonfig](https://github.com/crazy-max/gonfig)
* Add fields to load sensitive values from file
* Handle AWS EnvProvider for credentials
* Docs website with mkdocs
* Update deps

## 1.10.1 (2020/06/19)

* Fix unpublished Docker image

## 1.10.0 (2020/06/19)

* Add support for mips architectures (#160)
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.32.1
* Update deps

## 1.9.1 (2020/05/14)

* Typo

## 1.9.0 (2020/05/14)

* Allow using `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_HOSTED_ZONE_ID` environment variables (#130)
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

* Reenable multi-platform Docker images (#4)

## 1.3.1 (2019/08/31)

* Remove multi-platform Docker image, need rework (#4)

## 1.3.0 (2019/08/31)

* Multi-platform Docker image (#4)
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

* Run once on startup (#2)
* Update libs

## 0.2.0 (2019/04/02)

* Handle IPv4 or IPv6 only if included in a record set (#1)
* Update libs

## 0.1.1 (2019/04/02)

* Fix update of last IPv4/IPv6

## 0.1.0 (2019/04/01)

* Initial version based on [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.19.6
