# Multi stage build using go alpine image

# Stage 1
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything else
COPY . ./

# Generate
RUN go generate ./...

# Build the Go app
RUN go build -o tesetserver .

# Stage 2
FROM alpine:3.20

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/tesetserver /app/tesetserver

# Command to run the executable
CMD ["/app/tesetserver"]

# Build the docker image
# docker build -t testserver .
