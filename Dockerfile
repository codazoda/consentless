## Multi-stage build for consentless
#
# Build: docker build --tag consentless .
# Run:   docker run -p 8080:8080 -it --rm consentless

# 1) Builder: compile Go binary
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Cache modules (even though this project has no external deps)
COPY go.mod ./
RUN go mod download 2>/dev/null || true

# Copy source and build
COPY consentless.go ./
RUN go build -o consentless .

# 2) Runtime: copy the binary and run
FROM alpine:latest AS runtime
WORKDIR /app

# Optionally document the port
EXPOSE 8080

# Copy the statically linked binary from the builder
COPY --from=builder /app/consentless /usr/local/bin/consentless

# Default command
CMD ["consentless"]
