# syntax=docker/dockerfile:1

FROM golang

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy go files
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /cloudflare-ddns


# Run
CMD ["/cloudflare-ddns"]