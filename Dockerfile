# syntax=docker/dockerfile:1

ARG GO_VERSION="1.21"
ARG ALPINE_VERSION="3.18"
ARG XX_VERSION="1.3.0"

FROM --platform=$BUILDPLATFORM tonistiigi/xx:${XX_VERSION} AS xx

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
COPY --from=xx / /
ENV CGO_ENABLED=0
RUN apk add --no-cache file git
WORKDIR /src

FROM base AS version
ARG GIT_REF
RUN --mount=target=. <<EOT
  set -e
  case "$GIT_REF" in
    refs/tags/v*) version="${GIT_REF#refs/tags/}" ;;
    *) version=$(git describe --match 'v[0-9]*' --dirty='.m' --always --tags) ;;
  esac
  echo "$version" | tee /tmp/.version
EOT

FROM base AS vendored
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod download

FROM vendored AS test
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc linux-headers musl-dev
RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build <<EOT
  set -ex
  go test -v -coverprofile=/tmp/coverage.txt -covermode=atomic -race ./...
  go tool cover -func=/tmp/coverage.txt
EOT

FROM scratch AS test-coverage
COPY --from=test /tmp/coverage.txt /coverage.txt

FROM vendored AS build
ARG TARGETPLATFORM
RUN --mount=type=bind,target=. \
    --mount=type=bind,from=version,source=/tmp/.version,target=/tmp/.version \
    --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go/pkg/mod <<EOT
  set -ex
  xx-go build -trimpath -ldflags "-s -w -X main.version=$(cat /tmp/.version)" -o /usr/bin/ddns-route53 ./cmd
  xx-verify --static /usr/bin/ddns-route53
EOT

FROM scratch AS binary-unix
COPY --link --from=build /usr/bin/ddns-route53 /

FROM scratch AS binary-windows
COPY --link --from=build /usr/bin/ddns-route53 /ddns-route53.exe

FROM binary-unix AS binary-darwin
FROM binary-unix AS binary-freebsd
FROM binary-unix AS binary-linux
FROM binary-$TARGETOS AS binary
# enable scanning for this stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true

FROM --platform=$BUILDPLATFORM alpine:${ALPINE_VERSION} AS build-artifact
RUN apk add --no-cache bash tar zip
WORKDIR /work
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
RUN --mount=type=bind,target=/src \
    --mount=type=bind,from=binary,target=/build \
    --mount=type=bind,from=version,source=/tmp/.version,target=/tmp/.version <<EOT
  set -ex
  mkdir /out
  version=$(cat /tmp/.version)
  cp /build/* /src/CHANGELOG.md /src/LICENSE /src/README.md .
  if [ "$TARGETOS" = "windows" ]; then
    zip -r "/out/ddns-route53_${version#v}_${TARGETOS}_${TARGETARCH}${TARGETVARIANT}.zip" .
  else
    tar -czvf "/out/ddns-route53_${version#v}_${TARGETOS}_${TARGETARCH}${TARGETVARIANT}.tar.gz" .
  fi
EOT

FROM scratch AS artifact
COPY --link --from=build-artifact /out /

FROM scratch AS artifacts
FROM --platform=$BUILDPLATFORM alpine:${ALPINE_VERSION} AS releaser
RUN apk add --no-cache bash coreutils
WORKDIR /out
RUN --mount=from=artifacts,source=.,target=/artifacts <<EOT
  set -e
  cp /artifacts/**/* /out/ 2>/dev/null || cp /artifacts/* /out/
  sha256sum -b ddns-route53_* > ./checksums.txt
  sha256sum -c --strict checksums.txt
EOT

FROM scratch AS release
COPY --link --from=releaser /out /

FROM alpine:${ALPINE_VERSION}
RUN apk --update --no-cache add ca-certificates openssl shadow \
  && addgroup -g 1000 ddns-route53 \
  && adduser -u 1000 -G ddns-route53 -s /sbin/nologin -D ddns-route53
COPY --from=build /usr/bin/ddns-route53 /usr/local/bin/ddns-route53
USER ddns-route53
ENTRYPOINT [ "ddns-route53" ]
CMD [ "--config", "/ddns-route53.yml" ]
