# Build stage
FROM golang:alpine3.23 AS builder

WORKDIR /app

COPY go.sum go.mod .

RUN go mod download

COPY . .

ENV CGO_ENBABLED=0 GOOS=linux GOARCH=amd64

# Build
RUN go build -o main ./cmd/main.go

# Final image
FROM alpine:3.22

RUN addgroup -S srvgrp && adduser -S srvusr -G srvgrp

WORKDIR /server

COPY --from=builder /app/main .
COPY --from=builder /app/init_db_table/00_extensions.sql ./init_db_table/
COPY --from=builder /app/init_db_table/01_tables.sql ./init_db_table/

EXPOSE 3015

USER srvusr

ENTRYPOINT ["./main"]