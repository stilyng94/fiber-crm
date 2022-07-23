#!/usr/bin/env bash

set -e
# Show env vars
grep -v '^#' .env

# Export env vars
export $(grep -v '^#' .env | xargs)

migrate -source file://migrations -database "$DATABASE_URL" up

echo "done migration"

exec "$@"