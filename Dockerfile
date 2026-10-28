# --- Stage 1: Build the binary ---
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary statically, targeting Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Stage 2: Create the minimal runtime image ---
FROM alpine:3.20

# Add non-root user for security
RUN adduser -D appuser
USER appuser

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on
EXPOSE 8080

# Run the binary
CMD ["./main"]
