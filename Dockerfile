# syntax=docker/dockerfile:experimental
FROM --platform=amd64 golang:1.12.4 as builder

ARG TARGETPLATFORM
ARG VERSION

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go version
RUN go mod download
COPY . ./
RUN cp /usr/local/go/lib/time/zoneinfo.zip ./
RUN bash gobuild.sh ${TARGETPLATFORM} ${VERSION}

FROM --platform=$TARGETPLATFORM alpine:latest

LABEL maintainer="CrazyMax" \
  org.label-schema.name="ddns-route53" \
  org.label-schema.description="Dynamic DNS for Amazon Route 53‎ on a time-based schedule" \
  org.label-schema.url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vcs-url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vendor="CrazyMax" \
  org.label-schema.schema-version="1.0"

RUN uname -a \
  && apk --update --no-cache add \
    ca-certificates \
    libressl \
    tzdata \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=builder /app/ddns-route53 /usr/local/bin/ddns-route53
COPY --from=builder /app/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

ENTRYPOINT [ "ddns-route53", "--config", "/ddns-route53.yml" ]
