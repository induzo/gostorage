#!/usr/bin/env bash

set -euo pipefail

golangci-lint  run --no-config --issues-exit-code=0 --timeout=15s --disable-all \
    --enable govet \
    --enable errcheck \
    --enable staticcheck \
    --enable unused \
    --enable gosimple \
    --enable structcheck \
    --enable varcheck \
    --enable ineffassign \
    --enable deadcode \
    --enable typecheck \
    --enable bodyclose \
    --enable golint \
    --enable rowserrcheck \
    --enable stylecheck \
    --enable gosec \
    --enable interfacer \
    --enable unconvert \
    --enable goconst \
    --enable gocyclo \
    --enable gocognit \
    --enable gofmt \
    --enable goimports \
    --enable maligned \
    --enable depguard \
    --enable misspell \
    --enable lll \
    --enable unparam \
    --enable dogsled \
    --enable nakedret \
    --enable prealloc \
    --enable gocritic \
    --enable gochecknoinits \
    --enable gochecknoglobals \
    --enable godox \
    --enable funlen \
    --enable whitespace \
    --enable wsl \
    --enable goprintffuncname \
    --enable gomnd

    # --enable dupl \
    # --enable scopelint \

docker stack deploy --compose-file=infra/compose.yml storage;

# wait for the containers to be ready
sleep 20s

go test ./... -cover -bench=. -test.benchtime=3s -test.benchmem;

docker stack rm storage;