ARG GO_VERSION=1.15
ARG VERSION=dev

FROM --platform=${BUILDPLATFORM:-linux/amd64} tonistiigi/xx:golang AS xgo

FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:${GO_VERSION}-alpine AS base
RUN apk add --no-cache curl gcc git musl-dev
COPY --from=xgo / /
WORKDIR /src

FROM base AS gomod
COPY . .
RUN go mod download

FROM gomod AS build
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG VERSION
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.io,direct
RUN go build -ldflags "-w -s -X 'main.version=${VERSION}'" -v -o /opt/ddns-route53 cmd/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest
LABEL maintainer="CrazyMax"

RUN apk --update --no-cache add \
    ca-certificates \
    libressl \
    shadow \
    tzdata \
  && addgroup -g 1000 ddns-route53 \
  && adduser -u 1000 -G ddns-route53 -s /sbin/nologin -D ddns-route53 \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=build /opt/ddns-route53 /usr/local/bin/ddns-route53
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
RUN ddns-route53 --version

USER ddns-route53

ENTRYPOINT [ "/usr/local/bin/ddns-route53" ]
CMD [ "--config", "/ddns-route53.yml" ]
