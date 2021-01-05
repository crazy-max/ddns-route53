# syntax=docker/dockerfile:1.2
ARG GO_VERSION=1.15
ARG GORELEASER_VERSION=0.149.0

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS base
ARG GORELEASER_VERSION
RUN apk add --no-cache ca-certificates curl gcc file git musl-dev tar
RUN wget -qO- https://github.com/goreleaser/goreleaser/releases/download/v${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz | tar -zxvf - goreleaser \
  && mv goreleaser /usr/local/bin/goreleaser
WORKDIR /src

FROM base AS gomod
RUN --mount=type=bind,target=.,rw \
  --mount=type=cache,target=/go/pkg/mod \
  go mod tidy && go mod download

FROM gomod AS build
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
ARG GIT_REF
RUN --mount=type=bind,target=/src,rw \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=target=/go/pkg/mod,type=cache \
  ./hack/goreleaser.sh

FROM scratch AS artifacts
COPY --from=build /out/*.tar.gz /
COPY --from=build /out/*.zip /

FROM --platform=$TARGETPLATFORM alpine
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

ENTRYPOINT [ "ddns-route53" ]
CMD [ "--config", "/ddns-route53.yml" ]
