# Changelog

## 2.15.0 (2026/04/19)

* Add configurable WAN IP providers by @crazy-max in #1369
* Built-in WAN IP defaults now use AWS CheckIP and Cloudflare by @crazy-max in #1364
* Refactor WAN IP providers and client transport handling by @crazy-max in #1364
* Log startup config as structured debug field by @crazy-max in #1370
* Modernize config defaults and remove obsolete utl helpers by @crazy-max in #1366
* Document AWS shared credentials file support by @crazy-max in #1367
* Fix docs download links to use latest stable git tag by @crazy-max in #1363
* Go 1.26 by @crazy-max in #1361
* MkDocs Materials 9.7.5 by @crazy-max in #1362
* Bump github.com/alecthomas/kong to 1.15.0 in #1352
* Bump github.com/aws/aws-sdk-go-v2 dependencies in #1320 #1371
    * github.com/aws/aws-sdk-go-v2 to 1.41.5
    * github.com/aws/aws-sdk-go-v2/config to 1.32.15
    * github.com/aws/aws-sdk-go-v2/credentials to 1.19.14
    * github.com/aws/aws-sdk-go-v2/service/route53 to 1.62.5
* Bump github.com/crazy-max/gonfig to 0.8.0 in #1350
* Bump github.com/dromara/carbon/v2 to 2.6.16 in #1321
* Bump github.com/go-playground/validator/v10 to 10.30.2 in #1351
* Bump github.com/rs/zerolog to 1.35.0 in #1348
* Bump golang.org/x/sys to 0.43.0 in #1355

**Full Changelog**: [`v2.14.0...v2.15.0`](https://github.com/crazy-max/ddns-route53/compare/v2.14.0...v2.15.0)

## 2.14.0 (2025/12/30)

* Go 1.25 by @crazy-max in #1306
* Alpine Linux 3.23 by @crazy-max in #1309
* MkDocs Material 9.6.20 by @crazy-max in #1302 #1304
* Bump github.com/alecthomas/kong from 1.13.0 in #1318
* Bump github.com/aws/aws-sdk-go-v2 dependencies in #1292
    * github.com/aws/aws-sdk-go-v2 to 1.41.0
    * github.com/aws/aws-sdk-go-v2/config to 1.32.6
    * github.com/aws/aws-sdk-go-v2/credentials to 1.19.6
    * github.com/aws/aws-sdk-go-v2/service/route53 to 1.62.0
* Bump github.com/dromara/carbon/v2 to 2.6.15 in #1314
* Bump github.com/go-playground/validator/v10 to 10.30.1 in #1317
* Bump github.com/hashicorp/go-retryablehttp to 0.7.8 in #1315
* Bump github.com/rs/zerolog to 1.34.0 in #1296
* Bump github.com/stretchr/testify to 1.11.1 in #1311
* Bump golang.org/x/sys to 0.39.0 in #1293
* Bump golang.org/x/crypto to 0.45.0 in #1308

**Full Changelog**: [`v2.13.0...v2.14.0`](https://github.com/crazy-max/ddns-route53/compare/v2.13.0...v2.14.0)

## 2.13.0 (2024/12/20)

* Allow non-static AWS Credentials by @luna-lightblade in #1204
* Go 1.23 by @crazy-max in #1258
* Alpine Linux 3.21 by @crazy-max in #1258
* Switch to github.com/dromara/carbon/v2 2.5.2 by @crazy-max in #1259
* Bump github.com/alecthomas/kong to 1.6.0 in #1251
* Bump github.com/aws/aws-sdk-go-v2 dependencies in #1254
    * github.com/aws/aws-sdk-go-v2 to 1.32.7
    * github.com/aws/aws-sdk-go-v2/config to 1.28.7
    * github.com/aws/aws-sdk-go-v2/credentials to 1.17.48
    * github.com/aws/aws-sdk-go-v2/service/route53 to 1.46.4
* Bump github.com/go-playground/validator/v10 to 10.23.0 in #1219 #1253
* Bump github.com/hashicorp/go-retryablehttp to 0.7.7 in #1214
* Bump github.com/rs/zerolog to 1.33.0 in #1212
* Bump golang.org/x/crypto to 0.31.0 in #1257
* Bump golang.org/x/net to 0.23.0 in #1201
* Bump golang.org/x/sys to 0.28.0 in #1217 #1252

**Full Changelog**: [`v2.12.0...v2.13.0`](https://github.com/crazy-max/ddns-route53/compare/v2.12.0...v2.13.0)

## 2.12.0 (2024/03/24)

* Generate sbom and provenance by @crazy-max in #1193
* Bump github.com/alecthomas/kong to 0.9.0 in #1187
* Bump github.com/aws/aws-sdk-go-v2 dependencies in #1154 #1158 #1191
    * github.com/aws/aws-sdk-go-v2 to 1.26.0
    * github.com/aws/aws-sdk-go-v2/config to 1.27.9
    * github.com/aws/aws-sdk-go-v2/credentials to 1.17.9
    * github.com/aws/aws-sdk-go-v2/service/route53 to 1.40.3
* Bump github.com/go-playground/validator/v10 to 10.19.0 in #1183
* Bump github.com/rs/zerolog to 1.32.0 in #1172
* Bump golang.org/x/crypto to 0.17.0 in #1153
* Bump golang.org/x/sys to 0.18.0 in #1184

**Full Changelog**: [`v2.11.0...v2.12.0`](https://github.com/crazy-max/ddns-route53/compare/v2.11.0...v2.12.0)

## 2.11.0 (2023/12/16)

* Go 1.21 by @crazy-max in #1152
* Bump github.com/alecthomas/kong to 0.8.1 in #1077
* Bump github.com/aws/aws-sdk-go-v2 to 1.24.0 in #1144
* Bump github.com/aws/aws-sdk-go-v2/config to 1.26.1 in #1147
* Bump github.com/aws/aws-sdk-go-v2/credentials to 1.16.12 in #1145
* Bump github.com/aws/aws-sdk-go-v2/service/route53 to 1.35.5 in #1146
* Bump github.com/go-playground/validator/v10 to 10.16.0 in #1096
* Bump github.com/hashicorp/go-retryablehttp to 0.7.5 in #1101
* Bump golang.org/x/net to 0.17.0 in #1076
* Bump golang.org/x/sys to 0.15.0 in #1124

**Full Changelog**: [`v2.10.0...v2.11.0`](https://github.com/crazy-max/ddns-route53/compare/v2.10.0...v2.11.0)

## 2.10.0 (2023/09/27)

* Align AWS request retries with `max-retries` flag and decrease backoff delay to 5s by @crazy-max in #1064
* Add max backoff delay option by @crazy-max in #1065
* Go 1.20 by @crazy-max in #995
* Alpine Linux 3.18 by @crazy-max in #996
* Bump github.com/alecthomas/kong to 0.8.0 in #994
* Bump github.com/aws/aws-sdk-go-v2 to 1.21.0 in #989 #1035
* Bump github.com/aws/aws-sdk-go-v2/config to 1.18.42 in #908 #991 #1047 #1062
* Bump github.com/aws/aws-sdk-go-v2/credentials to 1.13.37 in #992 #1046
* Bump github.com/aws/aws-sdk-go-v2/service/route53 to 1.29.5 in #907 #993 #1037
* Bump github.com/crazy-max/gonfig to 0.7.1 in #975
* Bump github.com/go-playground/validator/v10 to 10.15.4 in #984 #1020 #1056
* Bump github.com/hashicorp/go-retryablehttp to 0.7.4 in #905 #985
* Bump github.com/rs/zerolog to 1.31.0 in #918 #961 #1011 #1063
* Bump github.com/stretchr/testify to 1.8.4 in #937 #982
* Bump golang.org/x/crypto to 0.1.0 in #939
* Bump golang.org/x/sys to 0.12.0 in #906 #938 #953 #969 #986 #997 #1019 #1044
* Bump golang.org/x/text to 0.3.8 in #936

**Full Changelog**: [`v2.9.0...v2.10.0`](https://github.com/crazy-max/ddns-route53/compare/v2.9.0...v2.10.0)

## 2.9.0 (2022/12/31)

* Switch to retryable http client implementation by @crazy-max in #900
* Allow specifying custom interface by @crazy-max in #901
* Check record value with WAN IP and update accordingly by @crazy-max in #902
    * :warning: requires `route53:ListResourceRecordSets` authorization in IAM policy
* Update to AWS SDK v2 by @crazy-max in #899
* Go 1.19 by @crazy-max in #846 #750
* Alpine Linux 3.17 by @crazy-max in #898
* Enhance workflow by @crazy-max in #847
* Bump github.com/go-playground/validator/v10 from 10.10.0 to 10.11.1 in #667 #710 #825
* Bump github.com/alecthomas/kong from 0.3.0 to 0.7.1 in #673 #749 #754 #864
* Bump github.com/stretchr/testify from 1.7.0 to 1.8.1 in #675 #743 #767 #848
* Bump github.com/rs/zerolog from 1.26.1 to 1.27.0 in #744
* Bump golang.org/x/sys to v0.3.0 in #751 #903
* Bump github.com/crazy-max/gonfig from 0.5.0 to 0.6.0 in #780
* Bump github.com/rs/zerolog from 1.27.0 to 1.28.0 in #810

**Full Changelog**: [`v2.8.0...v2.9.0`](https://github.com/crazy-max/ddns-route53/compare/v2.8.0...v2.9.0)

## 2.8.0 (2022/01/26)

* Alpine Linux 3.15 by @crazy-max in #622
* MkDocs Material 8.1.8 by @crazy-max in #637
* goreleaser-xx 1.2.5 by @crazy-max in #627
* Move `syscall` to `golang.org/x/sys` by @crazy-max in #612
* Move from `io/ioutil` to `io` and `os` packages by @crazy-max in #613
* Enhance dockerfiles by @crazy-max in #610 #611
* Bump github.com/go-playground/validator/v10 from 10.9.0 to 10.10.0 in #614
* Bump github.com/alecthomas/kong from 0.2.17 to 0.3.0 in #576 #599 #605 #620
* Bump github.com/aws/aws-sdk-go from 1.40.37 to 1.42.41 in #551 #577 #615 #623 #636
* Bump github.com/rs/zerolog from 1.24.0 to 1.26.1 in #534 #574 #607

**Full Changelog**: [`v2.7.0...v2.8.0`](https://github.com/crazy-max/ddns-route53/compare/v2.7.0...v2.8.0)

## 2.7.0 (2021/09/05)

* Go 1.17 by @crazy-max in #519
* Add `windows/arm64` artifact
* Wrong remaining time displayed
* Bump github.com/aws/aws-sdk-go from 1.38.69 to 1.40.37 in #518 #531
* Bump github.com/rs/zerolog from 1.23.0 to 1.24.0 in #525
* Bump github.com/crazy-max/gonfig from 0.4.0 to 0.5.0 in #520
* Bump codecov/codecov-action from 1 to 2
* Bump github.com/go-playground/validator/v10 from 10.6.1 to 10.9.0 in #483 #508
* Corrected Unifi documentation by @mhriemers in #477

**Full Changelog**: [`v2.6.1...v2.7.0`](https://github.com/crazy-max/ddns-route53/compare/v2.6.1...v2.7.0)

## 2.6.1 (2021/06/29)

* `windows/arm64` not yet available

**Full Changelog**: [`v2.6.0...v2.6.1`](https://github.com/crazy-max/ddns-route53/compare/v2.6.0...v2.6.1)

## 2.6.0 (2021/06/29)

* Bump github.com/aws/aws-sdk-go from 1.38.55 to 1.38.69 in #472
* Add `linux/riscv64` and `windows/arm64` artifacts by @crazy-max in #475
* Alpine Linux 3.14 by @crazy-max in #474
* Added quotations to linux service by @chriscn in #473
* Bump github.com/rs/zerolog from 1.22.0 to 1.23.0 in #463
* Bump github.com/alecthomas/kong from 0.2.16 to 0.2.17 in #457

**Full Changelog**: [`v2.5.0...v2.6.0`](https://github.com/crazy-max/ddns-route53/compare/v2.5.0...v2.6.0)

## 2.5.0 (2021/06/05)

* Allow disabling log color output by @crazy-max in #456
* Update build workflow by @crazy-max in #455
* Add CNI installation to unifi-os install by @amdprophet in #438
* MkDocs Materials 7.1.5 by @crazy-max in #453
* Fix artifacts download links
* Set `cacheonly` output for validators
* Move to `docker/metadata-action`
* Add `darwin/arm64` artifact
* Deploy docs on workflow dispatch or tag
* Go 1.16 by @crazy-max in #390
* Bump github.com/alecthomas/kong from 0.2.15 to 0.2.16 in #376
* Bump github.com/aws/aws-sdk-go from 1.37.25 to 1.38.55 in #377 #391 #454
* Bump github.com/go-playground/validator/v10 from 10.4.1 to 10.6.1 in #435
* Bump github.com/rs/zerolog from 1.20.0 to 1.22.0 in #388 #437

**Full Changelog**: [`v2.4.0...v2.5.0`](https://github.com/crazy-max/ddns-route53/compare/v2.4.0...v2.5.0)

## 2.4.0 (2021/03/07)

* Fix WanIP provider
* Missing non-root user for Docker image
* Switch to `goreleaser-xx` by @crazy-max in #373
* Remove s390x Docker image support by @crazy-max in #361
* Bump github.com/alecthomas/kong from 0.2.12 to 0.2.15 in #353
* Bump github.com/aws/aws-sdk-go from 1.36.20 to 1.37.25 in #374
* Bump github.com/stretchr/testify from 1.6.1 to 1.7.0 in #339

**Full Changelog**: [`v2.3.0...v2.4.0`](https://github.com/crazy-max/ddns-route53/compare/v2.3.0...v2.4.0)

## 2.3.0 (2021/01/05)

* Refactor CI and dev workflow with buildx bake by @crazy-max in #333
      * Upload artifacts
      * Add `image-local` target
      * Single job for artifacts and image
      * Add `armv5`, `ppc64le` and `s390x` artifacts
* Handle multi IP providers by @crazy-max in #331
* Bump github.com/alecthomas/kong from 0.2.11 to 0.2.12 in #307
* Bump github.com/aws/aws-sdk-go from 1.35.28 to 1.36.20 in #332

**Full Changelog**: [`v2.2.0...v2.3.0`](https://github.com/crazy-max/ddns-route53/compare/v2.2.0...v2.3.0)

## 2.2.0 (2020/11/15)

* Use embedded tzdata package
* Remove `--timezone` flag
* Docker image also available on [GitHub Container Registry](https://github.com/users/crazy-max/packages/container/package/ddns-route53)
* Switch to Docker actions
* Update deps

**Full Changelog**: [`v2.1.0...v2.2.0`](https://github.com/crazy-max/ddns-route53/compare/v2.1.0...v2.2.0)

## 2.1.0 (2020/09/07)

* Don't fill record change if IP address not available in #224
* Go 1.15
* Update deps

**Full Changelog**: [`v2.0.1...v2.1.0`](https://github.com/crazy-max/ddns-route53/compare/v2.0.1...v2.1.0)

## 2.0.1 (2020/08/05)

* Fix nil pointer with AWS credentials in #186

**Full Changelog**: [`v2.0.0...v2.0.1`](https://github.com/crazy-max/ddns-route53/compare/v2.0.0...v2.0.1)

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

**Full Changelog**: [`v1.10.1...v2.0.0`](https://github.com/crazy-max/ddns-route53/compare/v1.10.1...v2.0.0)

## 1.10.1 (2020/06/19)

* Fix unpublished Docker image

**Full Changelog**: [`v1.10.0...v1.10.1`](https://github.com/crazy-max/ddns-route53/compare/v1.10.0...v1.10.1)

## 1.10.0 (2020/06/19)

* Add support for mips architectures by @amdprophet in #160
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.32.1
* Update deps

**Full Changelog**: [`v1.9.1...v1.10.0`](https://github.com/crazy-max/ddns-route53/compare/v1.9.1...v1.10.0)

## 1.9.1 (2020/05/14)

* Typo

**Full Changelog**: [`v1.9.0...v1.9.1`](https://github.com/crazy-max/ddns-route53/compare/v1.9.0...v1.9.1)

## 1.9.0 (2020/05/14)

* Allow using `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_HOSTED_ZONE_ID` environment variables in #130
* Check AWS Route53 hosted zone ID entry
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.30.27

**Full Changelog**: [`v1.8.0...v1.9.0`](https://github.com/crazy-max/ddns-route53/compare/v1.8.0...v1.9.0)

## 1.8.0 (2020/05/06)

* Add `--log-caller` flag
* Flag `--log-json` not handled
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.30.20
* Update deps

**Full Changelog**: [`v1.7.0...v1.8.0`](https://github.com/crazy-max/ddns-route53/compare/v1.7.0...v1.8.0)

## 1.7.0 (2020/04/06)

* Switch to kong command-line parser
* Use Open Container Specification labels as label-schema.org ones are deprecated
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.30.3
* Update deps

**Full Changelog**: [`v1.6.0...v1.7.0`](https://github.com/crazy-max/ddns-route53/compare/v1.6.0...v1.7.0)

## 1.6.0 (2019/12/20)

* Strengthen WAN IP address retrieval and validation
* Move ident.me client to pkg
* Add `--max-retries` flag
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.26.6
* Update deps
* Go 1.13.5

**Full Changelog**: [`v1.5.0...v1.6.0`](https://github.com/crazy-max/ddns-route53/compare/v1.5.0...v1.6.0)

## 1.5.0 (2019/11/11)

* Seconds field is now optional
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.25.31
* Cache go modules

**Full Changelog**: [`v1.4.2...v1.5.0`](https://github.com/crazy-max/ddns-route53/compare/v1.4.2...v1.5.0)

## 1.4.2 (2019/10/08)

* Allow both IPv4 and IPv6 records by @ubipo in #7
* Use Docker buildx action v1

**Full Changelog**: [`v1.4.1...v1.4.2`](https://github.com/crazy-max/ddns-route53/compare/v1.4.1...v1.4.2)

## 1.4.1 (2019/09/27)

* Build on respective platforms
* Update deps
* Go 1.12.10

**Full Changelog**: [`v1.4.0...v1.4.1`](https://github.com/crazy-max/ddns-route53/compare/v1.4.0...v1.4.1)

## 1.4.0 (2019/09/26)

* Run Docker container as non-root user

**Full Changelog**: [`v1.3.4...v1.4.0`](https://github.com/crazy-max/ddns-route53/compare/v1.3.4...v1.4.0)

## 1.3.4 (2019/09/22)

* Use GOPROXY
* Update workflow
    * Use softprops/action-gh-release to create GitHub release
    * Use [ghaction-goreleaser](https://github.com/crazy-max/ghaction-goreleaser) GitHub Action
    * Use [ghaction-docker-buildx](https://github.com/crazy-max/ghaction-docker-buildx) GitHub Action
* Stop publishing Docker image on Quay

**Full Changelog**: [`v1.3.3...v1.3.4`](https://github.com/crazy-max/ddns-route53/compare/v1.3.3...v1.3.4)

## 1.3.3 (2019/09/08)

* Fix DockerHub latest tag

**Full Changelog**: [`v1.3.2...v1.3.3`](https://github.com/crazy-max/ddns-route53/compare/v1.3.2...v1.3.3)

## 1.3.2 (2019/08/31)

* Reenable multi-platform Docker images in #4

**Full Changelog**: [`v1.3.1...v1.3.2`](https://github.com/crazy-max/ddns-route53/compare/v1.3.1...v1.3.2)

## 1.3.1 (2019/08/31)

* Remove multi-platform Docker image, need rework in #4

**Full Changelog**: [`v1.3.0...v1.3.1`](https://github.com/crazy-max/ddns-route53/compare/v1.3.0...v1.3.1)

## 1.3.0 (2019/08/31)

* Multi-platform Docker image in #4
* Remove GitHub Package support

**Full Changelog**: [`v1.2.1...v1.3.0`](https://github.com/crazy-max/ddns-route53/compare/v1.2.1...v1.3.0)

## 1.2.1 (2019/08/29)

* Push Docker image to GitHub Package

**Full Changelog**: [`v1.2.0...v1.2.1`](https://github.com/crazy-max/ddns-route53/compare/v1.2.0...v1.2.1)

## 1.2.0 (2019/08/29)

* Switch to GitHub Actions

**Full Changelog**: [`v1.1.0...v1.2.0`](https://github.com/crazy-max/ddns-route53/compare/v1.1.0...v1.2.0)

## 1.1.0 (2019/07/18)

* Display next execution time
* Use v3 robfig/cron
* Go 1.12.4

**Full Changelog**: [`v1.0.0...v1.1.0`](https://github.com/crazy-max/ddns-route53/compare/v1.0.0...v1.1.0)

## 1.0.0 (2019/06/03)

* Run once on startup in #2
* Update libs

**Full Changelog**: [`v0.2.0...v1.0.0`](https://github.com/crazy-max/ddns-route53/compare/v0.2.0...v1.0.0)

## 0.2.0 (2019/04/02)

* Handle IPv4 or IPv6 only if included in a record set in #1
* Update libs

**Full Changelog**: [`v0.1.1...v0.2.0`](https://github.com/crazy-max/ddns-route53/compare/v0.1.1...v0.2.0)

## 0.1.1 (2019/04/02)

* Fix update of last IPv4/IPv6

**Full Changelog**: [`v0.1.0...v0.1.1`](https://github.com/crazy-max/ddns-route53/compare/v0.1.0...v0.1.1)

## 0.1.0 (2019/04/01)

* Initial version based on [aws-sdk-go](https://github.com/aws/aws-sdk-go) 1.19.6
