# Build stage
FROM golang:alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build
RUN go build -o main ./cmd/main.go

# Final image
FROM alpine:3.22

RUN apk add --no-cache curl

RUN addgroup -S srvgrp && adduser -S srvusr -G srvgrp


WORKDIR /server

COPY --from=builder /app/main .
COPY --from=builder /app/init_db_table/00_extensions.sql ./init_db_table/
COPY --from=builder /app/init_db_table/01_tables.sql ./init_db_table/

EXPOSE 3015

HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:3015/health ||  exit 1

USER srvusr

ENTRYPOINT ["./main"]