# Define the base image to use for building the application (Go 1.21 on Alpine Linux)
FROM golang:1.21-alpine as builder

# RUN openssl s_client -showcerts -connect proxy.golang.org:443 -servername proxy.golang.org < /dev/null 2>/dev/null | openssl x509 -outform PEM > /usr/local/share/ca-certificates/ca.crt

# Set permissions for the saved SSL certificate and update the system's certificate authorities
# RUN chmod 644 /usr/local/share/ca-certificates/ca.crt && update-ca-certificates

# Set the working directory within the builder container
WORKDIR /build

# Copy the source code into the builder container
COPY . /build

# Build the Go application for Linux (amd64) with CGO enabled
RUN CGO_ENABLED=0 GOOS=linux go build -o store-back .

# Create a new stage for the final application image (based on Alpine Linux)
FROM alpine:3.18 as hoster

# Copy configuration files, assets, templates and the built application from the builder stage
COPY --from=builder /build/.env ./.env
COPY --from=builder /build/store-back ./store-back
COPY --from=builder /build/migrations ./migrations

# Define the entry point for the final application image
ENTRYPOINT [ "./store-back" ]