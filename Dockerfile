# Use the official Golang image as the base image
FROM golang:1.21.6-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Copy the games.json file
COPY games.json .

# Build the Go application
RUN go build -o main

# Expose the port your application listens on (if applicable)
EXPOSE 8080

# Command to run your application
CMD ["./main"]
