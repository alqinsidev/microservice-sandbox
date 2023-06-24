# Start with a minimal base image for ARM64 architecture
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and cache the Go module dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o /app/auth-service

# Start with a clean, minimal base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/auth-service .

# Expose the port that the application listens on
EXPOSE 8080

# Set the command to run the application
CMD ["./auth-service"]
