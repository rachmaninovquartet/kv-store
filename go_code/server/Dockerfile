FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install git and ca-certificates (needed for go mod download)
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8000

# Set environment variables
ENV STORAGE_TYPE=redis
ENV REDIS_ADDR=redis:6379

# Run the application
CMD ["./main"] 