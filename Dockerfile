# Build stage
FROM golang:1.24.2 AS builder

# Set the Current Working Directory inside the container
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (important!)
COPY . .

# Build
RUN go build -o simple-crud-go ./cmd/api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["./simple-crud-go"]

