# Order Processing System API

Backend service for product and order management using Go, Echo, PostgreSQL, golang-migrate, and go-playground validator.

## Prerequisites

### Option A - Run with Docker (recommended)

- Docker
- Docker Compose

### Option B - Run without Docker

- Go 1.26.3+
- PostgreSQL 16+
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [swag CLI](https://github.com/swaggo/swag)

## Environment Setup

1. Copy env file:
   ```bash
   cp .env.example .env
   ```
2. Adjust `DATABASE_URL` if needed.

## Run with Docker

1. Start everything:
   ```bash
   docker compose up --build
   ```
2. Wait until API logs show `Starting API...`.
3. Verify endpoints:
   - Health: `http://localhost:8080/health`
   - Docs UI: `http://localhost:8080/docs`

Docker startup flow already does:

- wait for PostgreSQL readiness
- run migrations (`migrate up`)
- regenerate OpenAPI docs (`swag init`)
- start API server

## Run without Docker

1. Start PostgreSQL locally and ensure database exists:
   - database: `project_test_balabali`
   - user/password: adjust to your local setup
2. Copy env file and set local `DATABASE_URL` in `.env`.
3. Run migrations:
   ```bash
   migrate -path ./migrations -database "$DATABASE_URL" up
   ```
4. Generate OpenAPI docs:
   ```bash
   swag init -g ./cmd/api/main.go -o ./docs
   ```
5. Run API:
   ```bash
   go run ./cmd/api
   ```
6. Verify:
   - Health: `http://localhost:8080/health`
   - Swagger UI: `http://localhost:8080/swagger/index.html`

## Useful Commands

- Stop containers:
  ```bash
  docker compose down
  ```
- Stop containers and remove DB volume:
  ```bash
  docker compose down -v
  ```
