# syntax=docker/dockerfile:1.3

ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS base
RUN apk add --no-cache gcc musl-dev
WORKDIR /src

FROM base AS vendored
RUN --mount=type=bind,target=.,rw \
  --mount=type=cache,target=/go/pkg/mod \
  go mod tidy && go mod download

FROM vendored AS test
RUN --mount=type=bind,target=. \
  --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache \
  go test -v -coverprofile=/tmp/coverage.txt -covermode=atomic -race ./...

FROM scratch AS test-coverage
COPY --from=test /tmp/coverage.txt /coverage.txt
