FROM golang:1.17-alpine AS BUILDER

# Set the Current Working Directory inside the container
WORKDIR /undirect

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Install basic packages
RUN apk add \
    gcc \
    g++

# Download all the dependencies
RUN go get

# Build image
RUN go build .

FROM alpine:latest AS RUNNER

COPY --from=builder /undirect/undirect .

# This container exposes port 8080 to the outside world
EXPOSE 5000

# Run the executable
CMD ["./undirect"]