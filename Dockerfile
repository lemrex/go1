# Build stage
FROM golang:1.22-alpine AS builder

# Set the /app folder as the working directory
WORKDIR /app

# Copy all files from the current directory into the /app folder in the Docker image
COPY . /app/

# Build the application
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o api main.go

# Final stage
FROM scratch

# Copy the build result from the previous stage
COPY --from=builder /app/api /usr/local/bin/api

# Declare the command to be executed when running the container
ENTRYPOINT ["api"]
