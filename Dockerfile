# Start from the official Go image
FROM golang:1.25.1-alpine3.21 AS builder

WORKDIR /app

# Copy go.mod first for caching dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o server ./cmd

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]