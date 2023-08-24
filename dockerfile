# Start from the Go base image
FROM golang:1.21.0 AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the binary
RUN go build -o voter-api

# Use a minimal image to run the application
FROM debian:buster-slim

# Copy the binary
COPY --from=build /app/voter-api /voter-api

# Set the binary as the entry point of the container
ENTRYPOINT ["/voter-api"]

# Service listens on port 8080
EXPOSE 8080

