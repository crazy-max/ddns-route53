#!/usr/bin/env bash
set -e

TARGETPLATFORM=${1:-linux/amd64}
VERSION=${2:-snapshot}

case "$TARGETPLATFORM" in
  "linux/amd64")
    GOOS=linux
    GOARCH=amd64
    GOARM=
    ;;
  "linux/arm/v6")
    GOOS=linux
    GOARCH=arm
    GOARM=6
    ;;
  "linux/arm/v7")
    GOOS=linux
    GOARCH=arm
    GOARM=7
    ;;
  "linux/arm64")
    GOOS=linux
    GOARCH=arm64
    GOARM=
    ;;
  "linux/386")
    GOOS=linux
    GOARCH=386
    GOARM=
    ;;
  "linux/ppc64le")
    GOOS=linux
    GOARCH=ppc64le
    GOARM=
    ;;
  "linux/s390x")
    GOOS=linux
    GOARCH=s390x
    GOARM=
    ;;
esac

echo "TARGETPLATFORM=${TARGETPLATFORM}"
echo "VERSION=${VERSION}"
echo "GOOS=${GOOS}"
echo "GOARCH=${GOARCH}"
echo "GOARM=${GOARM}"

go build -ldflags "-w -s -X 'main.version=${VERSION}'" -v -o ddns-route53 cmd/main.go
