FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOBIN=/out go install github.com/swaggo/swag/cmd/swag@v1.16.6
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/api ./cmd/api

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates bash curl postgresql-client

COPY --from=builder /out/api /app/api
COPY --from=builder /out/swag /usr/local/bin/swag
COPY migrations /app/migrations
COPY docs /app/docs
COPY internal /app/internal
COPY cmd /app/cmd
COPY go.mod /app/go.mod
COPY docker/entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh \
    && curl -fsSL https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
        | tar -xz -C /usr/local/bin \
    && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate \
    && chmod +x /usr/local/bin/migrate

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]
