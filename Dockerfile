# syntax=docker/dockerfile:1.2
ARG GO_VERSION=1.16

FROM --platform=$BUILDPLATFORM crazymax/goreleaser-xx:latest AS goreleaser-xx
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS base
COPY --from=goreleaser-xx / /
RUN apk add --no-cache ca-certificates gcc file git linux-headers musl-dev tar
WORKDIR /src

FROM base AS build
ARG TARGETPLATFORM
ARG GIT_REF
RUN --mount=type=bind,target=/src,rw \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=target=/go/pkg/mod,type=cache \
  goreleaser-xx --debug \
    --name "ddns-route53" \
    --dist "/out" \
    --hooks="go mod tidy" \
    --hooks="go mod download" \
    --main="./cmd/main.go" \
    --ldflags="-s -w -X 'main.version={{.Version}}'" \
    --files="CHANGELOG.md" \
    --files="LICENSE" \
    --files="README.md"

FROM scratch AS artifacts
COPY --from=build /out/*.tar.gz /
COPY --from=build /out/*.zip /

FROM alpine
LABEL maintainer="CrazyMax"

RUN apk --update --no-cache add \
    ca-certificates \
    libressl \
    shadow \
  && addgroup -g 1000 ddns-route53 \
  && adduser -u 1000 -G ddns-route53 -s /sbin/nologin -D ddns-route53 \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=build /usr/local/bin/ddns-route53 /usr/local/bin/ddns-route53
RUN ddns-route53 --version

USER ddns-route53

ENTRYPOINT [ "ddns-route53" ]
CMD [ "--config", "/ddns-route53.yml" ]
