#!/usr/bin/env bash

set -euo pipefail

if [[ -z "${DATABASE_URL:-}" ]]; then
  echo "DATABASE_URL is required"
  exit 1
fi

echo "Waiting for PostgreSQL..."
until pg_isready -d "${DATABASE_URL}" >/dev/null 2>&1; do
  sleep 1
done
echo "PostgreSQL is ready"

echo "Running database migrations..."
migrate -path /app/migrations -database "${DATABASE_URL}" up

echo "Generating OpenAPI docs..."
swag init -g cmd/api/main.go -o docs

echo "Starting API..."
exec /app/api
