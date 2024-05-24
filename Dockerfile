# Dockerfile

# Use an official Golang runtime as a parent image
FROM golang:1.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o converter ./cmd/converter

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/converter .

# Copy static files
COPY --from=builder /app/Front ./Front

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./converter"]
