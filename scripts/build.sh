#!/bin/bash

# https://gist.github.com/mohanpedala/1e2ff5661761d3abd0385e8223e16425
set -euo pipefail

docker build --no-cache -t credit-calc:$(git rev-parse --short HEAD) \
  --build-arg VERSION=$(git describe --tags --dirty --always) \
  --build-arg COMMIT=$(git rev-parse --short HEAD) \
  --build-arg BRANCH=$(git rev-parse --abbrev-ref HEAD) \
  --build-arg BUILT=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --build-arg BUILDER=$(whoami)@$(hostname) \
  .