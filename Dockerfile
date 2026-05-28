FROM golang:1.26.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOBIN=/out go install github.com/swaggo/swag/cmd/swag@v1.16.6
RUN GOBIN=/out go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/api ./cmd/api

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates bash curl postgresql-client

COPY --from=builder /out/api /app/api
COPY --from=builder /out/swag /usr/local/bin/swag
COPY --from=builder /out/migrate /usr/local/bin/migrate
COPY migrations /app/migrations
COPY docs /app/docs
COPY internal /app/internal
COPY cmd /app/cmd
COPY go.mod /app/go.mod
COPY docker/entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh \
    && chmod +x /usr/local/bin/migrate

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]
