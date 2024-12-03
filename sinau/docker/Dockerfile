# Step 1: Use official Golang image as base
FROM golang:1.22.3-alpine as builder

# Step 2: Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Step 3: Install necessary dependencies
RUN apk update && apk add --no-cache git

# Step 4: Set working directory inside container
WORKDIR /app

# Step 5: Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 6: Copy the entire project to the container
COPY . .

# Step 7: Build the Go application
RUN go build -o main ./cmd/main.go

# Step 8: Use a minimal image for production
FROM alpine:latest

# Step 9: Set working directory
WORKDIR /root/

# Step 10: Copy binary from builder stage
COPY --from=builder /app/main .

# Step 11: Add any additional required files (e.g., configuration) sample:
# COPY ./files ./files

# Step 12: Expose the application port
EXPOSE 8080

# Step 13: Run the application
CMD ["./main"]
