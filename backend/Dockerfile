# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server/main.go

# Run stage
FROM scratch

WORKDIR /

# Copy the pre-built binary and config
COPY --from=builder /app/app .
COPY --from=builder /app/.env.example .env

EXPOSE 8080

ENTRYPOINT ["./app"]
