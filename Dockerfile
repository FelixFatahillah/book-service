# Base image with Go
FROM golang:1.23-alpine AS base

# Set working directory
WORKDIR /app

# Install necessary tools for Go
RUN apk add --no-cache git

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Cache dependencies
RUN go mod download

# Builder stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy cached dependencies
COPY --from=base /go/pkg /go/pkg

# Copy application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application

# Final stage (minimal image)
FROM alpine:3.20 AS runner

# Install certificates (if needed)
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/application /app/application

# Expose ports
EXPOSE ${PORT_SERVER}
EXPOSE ${PORT_GRPC}

# Set entrypoint
ENTRYPOINT ["/app/application"]

