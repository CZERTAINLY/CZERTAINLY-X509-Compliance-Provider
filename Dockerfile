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

COPY docker /app/docker

#
# Run Stage
#
FROM alpine:3.15

# add non root user czertainly
RUN addgroup --system --gid 10001 czertainly && adduser --system --home /opt/czertainly --uid 10001 --ingroup czertainly czertainly

COPY --from=builder /app/docker /
COPY --from=builder /app /opt/czertainly

WORKDIR /opt/czertainly

USER 10001

ENTRYPOINT ["/opt/czertainly/entry.sh"]
