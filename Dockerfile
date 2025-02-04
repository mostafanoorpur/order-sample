# Step 1: Build the Golang app
FROM golang:1.23.0 AS builder

# Set working directory
WORKDIR /app

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.8.0

# Verify goose installation
RUN ls -l /go/bin/goose && /go/bin/goose --version

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the application binary
RUN go build -o main ./cmd/server/main.go

# Step 2: Create the production image
FROM debian:bookworm-slim

# Install necessary tools
RUN apt-get update && apt-get install -y ca-certificates bash && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Add /usr/local/bin to PATH
ENV PATH="/usr/local/bin:$PATH"

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
# Copy the goose binary from the builder stage
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Copy the migrations and configuration files
COPY ./migrations ./migrations
COPY ./config.example.yaml ./config.yaml

# Expose the application port
EXPOSE 4000

# Run migrations and start the application
ENTRYPOINT ["sh", "-c", "/usr/local/bin/goose -dir ./migrations postgres \"host=$POSTGRES_DB_HOST user=$POSTGRES_DB_USER dbname=$POSTGRES_DB_NAME sslmode=disable password=$POSTGRES_DB_PASSWORD\" up && ./main"]