#!/usr/bin/env bash
set -euo pipefail

teardown() {
  echo "Tearing down..."
  docker compose down --remove-orphans
}

# Trap EXIT, INT (Ctrl+C), and TERM (kill) signals
trap teardown EXIT INT TERM

# Run the test container with build
docker compose run --build test
