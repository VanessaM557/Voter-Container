# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.17 AS build

# Copy the local package files to the container's workspace.
WORKDIR /go/src/app
COPY . .

# Build the command inside the container.
RUN go get -d -v ./...
RUN go install -v ./...

# Use a Debian slim image for the runtime environment
FROM debian:buster-slim

# Copy the build artifact from the build stage.
COPY --from=build /go/bin/app /app

# Set the binary as the entry point of the container
ENTRYPOINT ["/app"]

# Service listens on port 8080
EXPOSE 8080
