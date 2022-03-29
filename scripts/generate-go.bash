#!/usr/bin/env bash
set -euo pipefail
: "${DEBUG=false}"
[[ "${DEBUG,,}" = 'true' ]] && set -vx

cd "$(dirname "$0")/.."

protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/simplegrpc.proto
