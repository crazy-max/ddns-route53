# syntax=docker/dockerfile:experimental
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.12-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN printf "I am running on ${BUILDPLATFORM:-linux/amd64}, building for ${TARGETPLATFORM:-linux/amd64}\n$(uname -a)\n"

RUN [ "$TARGETPLATFORM" = "linux/amd64"   ] && echo GOOS=linux GOARCH=amd64 > .env || true
RUN [ "$TARGETPLATFORM" = "linux/arm/v6"  ] && echo GOOS=linux GOARCH=arm GOARM=6 > .env || true
RUN [ "$TARGETPLATFORM" = "linux/arm/v7"  ] && echo GOOS=linux GOARCH=arm GOARM=7 > .env || true
RUN [ "$TARGETPLATFORM" = "linux/arm64"   ] && echo GOOS=linux GOARCH=arm64 > .env || true
RUN [ "$TARGETPLATFORM" = "linux/386"     ] && echo GOOS=linux GOARCH=386 > .env || true
RUN [ "$TARGETPLATFORM" = "linux/ppc64le" ] && echo GOOS=linux GOARCH=ppc64le > .env || true
RUN [ "$TARGETPLATFORM" = "linux/s390x"   ] && echo GOOS=linux GOARCH=s390x > .env || true
RUN env $(cat .env | xargs) go env

RUN apk --update --no-cache add \
    build-base \
    gcc \
    git \
  && rm -rf /tmp/* /var/cache/apk/*

WORKDIR /app

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
COPY go.mod .
COPY go.sum .
RUN env $(cat .env | xargs) go mod download
COPY . ./

ARG VERSION=dev
RUN env $(cat .env | xargs) go build -ldflags "-w -s -X 'main.version=${VERSION}'" -v -o ddns-route53 cmd/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest

LABEL maintainer="CrazyMax" \
  org.label-schema.name="ddns-route53" \
  org.label-schema.description="Dynamic DNS for Amazon Route 53‎ on a time-based schedule" \
  org.label-schema.url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vcs-url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vendor="CrazyMax" \
  org.label-schema.schema-version="1.0"

RUN apk --update --no-cache add \
    ca-certificates \
    libressl \
    shadow \
  && addgroup -g 1000 ddns-route53 \
  && adduser -u 1000 -G ddns-route53 -s /sbin/nologin -D ddns-route53 \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=builder /app/ddns-route53 /usr/local/bin/ddns-route53
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
RUN ddns-route53 --version

USER ddns-route53

ENTRYPOINT [ "ddns-route53" ]
CMD [ "--config", "/ddns-route53.yml" ]
