FROM golang:1.12.4 as builder

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go version
RUN go mod download
COPY . ./
RUN cp /usr/local/go/lib/time/zoneinfo.zip ./ \
  && CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-w -s -X 'main.version=${VERSION}'" \
    -v -o ddns-route53 cmd/main.go

FROM alpine:latest

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

LABEL maintainer="CrazyMax" \
  org.label-schema.build-date=$BUILD_DATE \
  org.label-schema.name="ddns-route53" \
  org.label-schema.description="Dynamic DNS for Amazon Route 53‎ on a time-based schedule" \
  org.label-schema.version=$VERSION \
  org.label-schema.url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vcs-ref=$VCS_REF \
  org.label-schema.vcs-url="https://github.com/crazy-max/ddns-route53" \
  org.label-schema.vendor="CrazyMax" \
  org.label-schema.schema-version="1.0"

RUN apk --update --no-cache add \
    ca-certificates \
    libressl \
    tzdata \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=builder /app/ddns-route53 /usr/local/bin/ddns-route53
COPY --from=builder /app/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

ENTRYPOINT [ "ddns-route53", "--config", "/ddns-route53.yml" ]
