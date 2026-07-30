# Use the official Golang image for building
FROM golang:1.26-alpine AS builder

# Set working directory
WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod go.sum ./

RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o web ./cmd/web

# Use a minimal base image for final deployment
FROM alpine:3

RUN addgroup -S app && adduser -S -G app app

# Set working directory in the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/web .

# Copy web UI assets
COPY --from=builder /app/ui ./ui

USER app

# Expose the application port
EXPOSE 4000

# Run the application
CMD ["./web"]
