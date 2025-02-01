# Use Golang base image for building
FROM golang:1.22 AS builder

# Set working directory
WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy application files
COPY . .

# Build the Go binary with Linux compatibility
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# Use a lightweight base image for the final container
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Install necessary dependencies (ca-certificates for HTTPS support)
RUN apk --no-cache add ca-certificates

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Expose API port
EXPOSE 8080

# Start the backend service
CMD ["./main"]
