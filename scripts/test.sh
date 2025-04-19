set -euo pipefail

teardown() {
  docker compose down --remove-orphans
}

TRAPINT() {
  teardown
}

trap teardown EXIT

docker compose run --build test
