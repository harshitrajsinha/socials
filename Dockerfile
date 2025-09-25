# Stage 1: Build
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o project ./cmd/main.go

# Stage 2: Runtime
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/project /app/project

EXPOSE 3015

CMD ["./project"]
