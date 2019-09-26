# syntax=docker/dockerfile:experimental
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.12.4-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN printf "I am running on ${BUILDPLATFORM:-linux/amd64}, building for ${TARGETPLATFORM:-linux/amd64}\n$(uname -a)\n"

RUN apk --update --no-cache add \
    bash \
    build-base \
    gcc \
    git \
  && rm -rf /tmp/* /var/cache/apk/*

WORKDIR /app

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
COPY go.mod .
COPY go.sum .
RUN go version
RUN go mod download
COPY . ./

ARG VERSION=dev
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
RUN bash gobuild.sh ${TARGETPLATFORM} ${VERSION}

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest

LABEL maintainer="CrazyMax" \
  org.label-schema.name="ddns-route53" \
  org.label-schema.description="Dynamic DNS for Amazon Route 53‎ on a time-based schedule" \
  org.label-schema.url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vcs-url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vendor="CrazyMax" \
  org.label-schema.schema-version="1.0"

ENV TZ="UTC"

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

ENTRYPOINT [ "ddns-route53" ]
CMD [ "--config", "/ddns-route53.yml" ]
