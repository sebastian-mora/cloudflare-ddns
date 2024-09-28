# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Install git (needed for Go module retrieval)
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o /cloudflare-ddns .

# Stage 2: Create a minimal image
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /cloudflare-ddns /usr/local/bin/cloudflare-ddns

# Run the binary
CMD ["/usr/local/bin/cloudflare-ddns"]
