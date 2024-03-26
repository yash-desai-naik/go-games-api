# Dockerfile

# Use the official Golang image as base
FROM golang:1.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules dependency files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a minimal base image for the final runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
