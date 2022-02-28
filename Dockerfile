# syntax=docker/dockerfile:1

# Start from go base image
FROM golang:1.16-alpine

# Add maintainer info
LABEL maintainer="Isuru Siriwardana <isurusiri@protonmail.com>"

# Creates working directory
WORKDIR /app

# Copy module files inside working directory
COPY go.mod ./
COPY go.sum ./

# Download go modules
RUN go mod download

# Copy source code inside working directory
COPY *.go ./

# Compile source code
RUN go build -o /kvsync

# Execute built artifact
ENTRYPOINT [ "/kvsync" ]
