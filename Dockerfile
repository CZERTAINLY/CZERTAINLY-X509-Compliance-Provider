# syntax=docker/dockerfile
# Install golang

# Build Stage
FROM golang:1.18-alpine3.15 AS builder

ENV WRK_DIR /app

# Copy the contents to /app
COPY . $WRK_DIR

# Set working directory
WORKDIR $WRK_DIR

# Toggle CGO based on your app requirement. CGO_ENABLED=1 for enabling CGO
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o $WRK_DIR/appbin $WRK_DIR

#
# Run Stage
#
FROM alpine:3.15

ENV WRK_DIR /app
COPY --from=builder $WRK_DIR $WRK_DIR

WORKDIR $WRK_DIR

# Start the app
CMD ["./appbin"]