#!/usr/bin/env bash

set -euo pipefail

golangci-lint run

docker-compose up -d
# wait for the containers to be ready
sleep 20s

go test ./... -cover -bench=. -test.benchtime=3s -test.benchmem

docker-compose down
