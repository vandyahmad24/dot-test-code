# Stage 1: Build the Go application
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY .env .


RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# Stage 2: Create the runtime container
FROM alpine:latest
WORKDIR /app

# Install tzdata for timezone data
RUN apk add --no-cache tzdata

# Copy the built application and other necessary files from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/.env .

# Set the timezone to Asia/Jakarta
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone

# Expose the application port
EXPOSE 8000

# Start the application
CMD ["./app"]