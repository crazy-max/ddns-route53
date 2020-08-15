FROM --platform=${BUILDPLATFORM:-linux/amd64} tonistiigi/xx:golang AS xgo
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.13-alpine AS builder

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION=dev

ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io,direct
COPY --from=xgo / /

RUN apk --update --no-cache add \
    build-base \
    gcc \
    git \
  && rm -rf /tmp/* /var/cache/apk/*

WORKDIR /app

COPY . ./
RUN go mod download

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
RUN go env
RUN go build -ldflags "-w -s -X 'main.version=${VERSION}'" -v -o ddns-route53 cmd/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

LABEL maintainer="CrazyMax" \
  org.opencontainers.image.created=$BUILD_DATE \
  org.opencontainers.image.url="https://github.com/crazy-max/ddns-route53" \
  org.opencontainers.image.source="https://github.com/crazy-max/ddns-route53" \
  org.opencontainers.image.version=$VERSION \
  org.opencontainers.image.revision=$VCS_REF \
  org.opencontainers.image.vendor="CrazyMax" \
  org.opencontainers.image.title="ddns-route53" \
  org.opencontainers.image.description="Dynamic DNS for Amazon Route 53‎ on a time-based schedule" \
  org.opencontainers.image.licenses="MIT"

RUN apk --update --no-cache add \
    ca-certificates \
    libressl \
    shadow \
    tzdata \
  && addgroup -g 1000 ddns-route53 \
  && adduser -u 1000 -G ddns-route53 -s /sbin/nologin -D ddns-route53 \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=builder /app/ddns-route53 /usr/local/bin/ddns-route53
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
RUN ddns-route53 --version

USER ddns-route53

ENTRYPOINT [ "/usr/local/bin/ddns-route53" ]
CMD [ "--config", "/ddns-route53.yml" ]
